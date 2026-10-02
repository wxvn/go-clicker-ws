package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wxvn/go-clicker-ws/internal/config"
	authRepo "github.com/wxvn/go-clicker-ws/internal/feature/auth/repository"
	authServ "github.com/wxvn/go-clicker-ws/internal/feature/auth/service"
	authhttp "github.com/wxvn/go-clicker-ws/internal/feature/auth/transport/http"
	"github.com/wxvn/go-clicker-ws/internal/feature/clicker/repository"
	"github.com/wxvn/go-clicker-ws/internal/feature/clicker/service"
	"github.com/wxvn/go-clicker-ws/internal/feature/clicker/transport/http"
	"github.com/wxvn/go-clicker-ws/internal/feature/clicker/transport/ws"
	"github.com/wxvn/go-clicker-ws/internal/hasher"
	"github.com/wxvn/go-clicker-ws/internal/logger"
	"github.com/wxvn/go-clicker-ws/internal/mongodb"
	"github.com/wxvn/go-clicker-ws/internal/redis"
	"github.com/wxvn/go-clicker-ws/internal/session"
	httptransport "github.com/wxvn/go-clicker-ws/internal/transport/http"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Logger)
	log.Info("init logger")

	ctx := context.Background()

	initCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	db, err := mongodb.New(cfg.MongoDB)
	if err != nil {
		log.Error("failed to connect to MongoDB", "error", err)
		return
	}

	if err := db.Ping(initCtx); err != nil {
		log.Error("failed to ping MongoDB", "error", err)
		return
	}

	if err := db.Init(initCtx, cfg.MongoDB.Database); err != nil {
		log.Error("failed to init MongoDB", "error", err)
		return
	}

	log.Info("MongoDB connected")

	redisClient := redis.New(cfg.Redis)
	sessionStore := session.NewStore(redisClient, time.Hour*24*7)

	// Repository
	userRepository := authRepo.NewAuthRepository(
		db,
		cfg.MongoDB.Database,
	)

	// Hasher
	passwordHasher := hasher.NewBcrypt()

	// Service
	authService := authServ.NewAuthService(
		userRepository,
		passwordHasher,
		sessionStore,
	)

	// HTTP handler
	authHandler := authhttp.NewHandler(authService)

	clickRepository := repository.NewClickerRepository(db, cfg.MongoDB.Database)
	clickService := service.NewClickService(clickRepository)
	clickHandler := http.NewHandler(clickService)

	hub := ws.NewHub()
	go hub.Run()

	clickHandlerWs := ws.NewHandler(hub, clickService, cfg.AllowedOrigins)

	// HTTP router
	handler := httptransport.NewRouter(authHandler, clickHandler, clickHandlerWs, log, sessionStore, cfg.AllowedOrigins)

	// HTTP server
	addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	if err := httptransport.RunServer(ctx, addr, handler); err != nil {
		log.Error("HTTP server stopped", "error", err)
	}
}
