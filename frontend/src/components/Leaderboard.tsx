import type { LeaderboardResponse } from '../features/leaderboard/types'

interface LeaderboardProps {
  data: LeaderboardResponse | null
}

const medals: Record<number, string> = { 1: '🥇', 2: '🥈', 3: '🥉' }

const rowStyles: Record<number, string> = {
  1: 'bg-yellow-200/70',
  2: 'bg-gray-200/70',
  3: 'bg-orange-200/70',
}

function Leaderboard({ data }: LeaderboardProps) {
  if (!data) {
    return <p className="text-sm text-gray-600">Loading...</p>
  }

  const isCurrentUserInList = data.users.some((u) => u.id === data.current_user.id)

  return (
    <div className="flex flex-col gap-1.5">
      <div className="flex items-center gap-3 px-2 text-xs font-semibold uppercase text-gray-600">
        <span className="w-6 text-center">#</span>
        <span className="flex-1">Player</span>
        <span className="w-14 text-right">Clicks</span>
      </div>

      {data.users.map((u) => {
        const isMe = u.id === data.current_user.id
        return (
          <div
            key={u.id}
            className={`flex items-center gap-3 rounded-lg px-2 py-1.5 transition-colors ${
              isMe ? 'bg-mauve-400/70 ring-2 ring-mauve-600' : (rowStyles[u.position] ?? '')
            }`}
          >
            <span className="w-6 text-center text-base">{medals[u.position] ?? u.position}</span>
            <span className="flex-1 truncate text-sm font-semibold text-gray-900">{u.username}</span>
            <span className="w-14 text-right text-sm font-bold text-gray-900">{u.clicks}</span>
          </div>
        )
      })}

      {!isCurrentUserInList && (
        <div className="mt-2 flex items-center gap-3 rounded-lg border-t border-gray-400/50 bg-mauve-400/70 px-2 py-1.5 pt-3 ring-2 ring-mauve-600">
          <span className="w-6 text-center text-base">{data.current_user.position}</span>
          <span className="flex-1 truncate text-sm font-semibold text-gray-900">{data.current_user.username}</span>
          <span className="w-14 text-right text-sm font-bold text-gray-900">{data.current_user.clicks}</span>
        </div>
      )}
    </div>
  )
}

export default Leaderboard