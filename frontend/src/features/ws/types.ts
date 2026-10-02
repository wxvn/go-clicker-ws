export interface ClickMessage {
  action: 'click'
  data: { clicks: number }
}

export interface LeaderboardUser {
  id: string
  username: string
  avatar: number
  clicks: number
  position: number
}

export interface LeaderboardMessage {
  action: 'leaderboard'
  data: { users: LeaderboardUser[] }
}

export interface ErrorMessage {
  action: 'error'
  error: string
}

export type WSMessage = ClickMessage | LeaderboardMessage | ErrorMessage