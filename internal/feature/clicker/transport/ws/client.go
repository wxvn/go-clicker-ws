package ws

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	userID  string
	service Service
}

type Request struct {
	Action string `json:"action"`
	Limit  int    `json:"limit,omitempty"`
}

type Response struct {
	Action string `json:"action"`
	Data   any    `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

type ClickResponse struct {
	Clicks int64 `json:"clicks"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   int    `json:"avatar"`
	Clicks   int64  `json:"clicks"`
}

type LeaderboardUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   int    `json:"avatar"`
	Clicks   int64  `json:"clicks"`
	Position int64  `json:"position"`
}

type LeaderboardResponse struct {
	Users []LeaderboardUserResponse `json:"users"`
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var req Request
		if err := json.Unmarshal(message, &req); err != nil {
			log.Printf("invalid json: %v", err)
			continue
		}

		switch req.Action {

		case "click":
			clicks, err := c.service.IncrementClicks(context.Background(), c.userID)
			if err != nil {
				res := Response{
					Action: req.Action,
					Error:  err.Error(),
				}
				data, marshalErr := json.Marshal(res)
				if marshalErr != nil {
					continue
				}
				c.send <- data
				continue
			}
			clickRes := Response{
				Action: "click",
				Data: ClickResponse{
					Clicks: clicks,
				},
			}
			clickData, err := json.Marshal(clickRes)
			if err != nil {
				continue
			}
			c.send <- clickData

			leaderboard, err := c.service.GetLeaderboard(context.Background(), 10, "")
			if err != nil {
				log.Printf("get leaderboard after click: %v", err)
				continue
			}

			users := make([]LeaderboardUserResponse, 0, len(leaderboard.Users))
			for _, u := range leaderboard.Users {
				var position int64
				if u.Position != nil {
					position = *u.Position
				}
				users = append(users, LeaderboardUserResponse{
					ID:       u.ID,
					Username: u.Username,
					Avatar:   u.Avatar,
					Clicks:   u.Clicks,
					Position: position,
				})
			}

			res := Response{
				Action: "leaderboard",
				Data:   LeaderboardResponse{Users: users},
			}

			data, err := json.Marshal(res)
			if err != nil {
				continue
			}

			c.hub.broadcast <- data

		default:
			res := Response{
				Action: req.Action,
				Error:  "unknown action",
			}

			data, err := json.Marshal(res)
			if err != nil {
				continue
			}

			c.send <- data

		}

	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
