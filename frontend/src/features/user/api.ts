import type { User } from './types'

const BASE_URL = import.meta.env.VITE_API_URL

export async function getMe(): Promise<User> {
  const response = await fetch(`${BASE_URL}/user/me`, {
    credentials: 'include',
  })

  if (!response.ok) {
    throw new Error('not authenticated')
  }

  return response.json()
}