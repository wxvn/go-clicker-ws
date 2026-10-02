import type { LeaderboardResponse } from './types'

const BASE_URL = import.meta.env.VITE_API_URL

export async function getLeaderboard(limit: number = 10): Promise<LeaderboardResponse> {
  const response = await fetch(`${BASE_URL}/leaderboard?limit=${limit}`, {
    credentials: 'include',
  })

  if (!response.ok) {
    throw new Error('failed to load leaderboard')
  }

  return response.json()
}