import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { logout } from '../features/auth/api'
import { useUser } from '../features/user/UserContext'
import avatar from '../assets/img/profile/cat-avatar.png'

function UserMenu() {
  const { user, setUser } = useUser()
  const [isOpen, setIsOpen] = useState(false)
  const navigate = useNavigate()

  async function handleLogout() {
    try {
      await logout()
    } catch (err) {
    }
    setUser(null)
    navigate('/login')
  }

  return (
    <div className="relative">
      <button
        onClick={() => setIsOpen((prev) => !prev)}
        className="rounded-full ring-2 ring-white/20 transition hover:ring-white/40 active:scale-95"
      >
        <img
          src={avatar}
          alt=""
          className="h-16 w-16 rounded-full object-cover shadow-md"
        />
      </button>

      {isOpen && (
        <div className="absolute left-1/2 top-full z-10 mt-2 w-36 -translate-x-1/2 rounded-lg border border-white/15 bg-zinc-800 p-1 shadow-lg">
          {user ? (
            <button
              onClick={handleLogout}
              className="w-full rounded-md px-3 py-2 text-left text-sm text-red-400 transition hover:bg-white/10"
            >
              Log out
            </button>
          ) : (
            <>
              <Link
                to="/login"
                className="block rounded-md px-3 py-2 text-left text-sm text-white transition hover:bg-white/10"
              >
                Login
              </Link>
              <Link
                to="/register"
                className="block rounded-md px-3 py-2 text-left text-sm text-white transition hover:bg-white/10"
              >
                Register
              </Link>
            </>
          )}
        </div>
      )}
    </div>
  )
}

export default UserMenu