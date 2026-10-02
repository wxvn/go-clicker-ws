const BASE_URL = import.meta.env.VITE_API_URL

interface ClickResponse {
  clicks: number
}

export async function click(): Promise<ClickResponse> {
  const response = await fetch(`${BASE_URL}/click`, {
    method: 'POST',
    credentials: 'include',
  })

  if (!response.ok) {
    throw new Error('failed to register click')
  }

  return response.json()
}