package domain

type Leaderboard struct {
	ID       string
	Username string
	Avatar   int
	Clicks   int64
	Position *int64
}

type LeaderboardResult struct {
	Users       []Leaderboard
	CurrentUser *Leaderboard
}
