export interface LeaderboardEntry {
  id: string
  username: string
  avatar: number
  clicks: number
  position: number
}

export interface LeaderboardResponse {
  users: LeaderboardEntry[]
  current_user: LeaderboardEntry
}