import { useEffect, useRef, useState } from 'react'
import { useUser } from '../features/user/UserContext'
import { getLeaderboard } from '../features/leaderboard/api'
import type { LeaderboardResponse } from '../features/leaderboard/types'
import { createSocket } from '../features/ws/socket'
import Leaderboard from '../components/Leaderboard'
import UserMenu from '../components/UserMenu'

function HomePage() {
  const { user } = useUser()
  const [clicks, setClicks] = useState(0)
  const [leaderboard, setLeaderboard] = useState<LeaderboardResponse | null>(null)
  const [isConnected, setIsConnected] = useState(false)
  const socketRef = useRef<ReturnType<typeof createSocket> | null>(null)

  useEffect(() => {
    if (user) setClicks(user.clicks)
  }, [user])

  useEffect(() => {
    function loadLeaderboard() {
      getLeaderboard(10)
        .then((res) => setLeaderboard(res))
        .catch(() => { })
    }

    loadLeaderboard()
    const interval = setInterval(loadLeaderboard, 15000)
    return () => clearInterval(interval)
  }, [])

  useEffect(() => {
    const socket = createSocket((message) => {
      if (message.action === 'click') {
        setClicks(message.data.clicks)

        setLeaderboard((prev) => {
          if (!prev) return prev
          return {
            ...prev,
            current_user: { ...prev.current_user, clicks: message.data.clicks },
          }
        })
      }

      if (message.action === 'leaderboard') {
        setLeaderboard((prev) => {
          if (!prev) return prev
          return { ...prev, users: message.data.users }
        })
      }
    })

    socket.socket.onopen = () => setIsConnected(true)
    socket.socket.onclose = () => setIsConnected(false)

    socketRef.current = socket

    return () => socket.close()
  }, [])

  function handleClick() {
    socketRef.current?.sendClick()
  }

  return (
    <div className="grid min-h-screen grid-cols-[1fr_auto_1fr] items-center bg-linear-to-br from-black to-gray-700">
      <div></div>

      <div className="flex flex-col items-center gap-6">
        <h1 className="text-5xl font-bold text-zinc-300">Clicker</h1>

        <div className="flex gap-6 rounded-xl border border-white/15 bg-white/10 p-6 shadow-lg backdrop-blur">

          <div className="flex w-56 flex-col items-center justify-between gap-8 py-2">
            <div className="flex flex-col items-center gap-2 text-center">
              <UserMenu />
              <p className="text-xl font-bold text-white">{user?.username ?? 'Guest'}</p>
              <p className="text-sm text-zinc-400">{clicks} clicks</p>
            </div>

            <button
              onClick={handleClick}
              disabled={!isConnected}
              className="h-44 w-44 rounded-full bg-red-500 text-3xl font-mono uppercase text-white shadow-lg transition active:scale-95 disabled:cursor-not-allowed disabled:opacity-50"
            >
              Press!
            </button>
          </div>

          <div className="flex w-64 flex-col rounded-xl bg-linear-to-b from-zinc-200 to-zinc-400 p-4 shadow-inner">
            <h2 className="mb-3 text-xl font-bold text-gray-900">Leaderboard</h2>
            <Leaderboard data={leaderboard} />
          </div>

        </div>
      </div>

      <div></div>
    </div>
  )
}

export default HomePage