import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { login } from '../features/auth/api'
import { useUser } from '../features/user/UserContext'
import Toast from '../components/ui/Toast'
import eyeIcon from '../assets/img/eye.svg'
import eyeOffIcon from '../assets/img/eye-off.svg'


function LoginPage() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [showPassword, setShowPassword] = useState(false)
  const navigate = useNavigate()
  const [error, setError] = useState('')
  const { setUser } = useUser()

  async function handleLogin() {
    setError('')
    try {
      const result = await login({ username, password })
      setUser(result)
      navigate('/')
    } catch (err) {
      setError('Login failded')
    }
  }

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-linear-to-br from-black to-gray-700">
      <Toast message={error} onClose={() => setError('')} />

      <h1 className="mb-4 text-4xl font-bold text-zinc-300">Login</h1>

      <div className="flex w-full max-w-sm flex-col items-center gap-1 rounded-xl border border-white/15 bg-white/10 p-10 shadow-lg backdrop-blur">

        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="my-1.5 w-full rounded-md border-2 bg-gray-200 p-1 outline-none"
        />

        <div className="relative my-1.5 w-full">
          <input
            type={showPassword ? 'text' : 'password'}
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full rounded-md border-2 bg-gray-200 p-1 pr-10 outline-none"
          />
          <button
            type="button"
            onClick={() => setShowPassword((prev) => !prev)}
            className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-600 hover:text-gray-900"
          >
            <img
              src={showPassword ? eyeOffIcon : eyeIcon}
              alt={showPassword ? 'Hide password' : 'Show password'}
              className="h-5 w-5"
            />
          </button>
        </div>

        <button
          onClick={handleLogin}
          className="mt-6 w-1/2 rounded-md bg-zinc-500 p-2 font-bold text-gray-800 transition active:scale-95 hover:bg-zinc-600">
          login
        </button>

        <p className="mt-4 text-sm text-zinc-400">
          Don't have an account?{' '}
          <Link to="/register" className="text-white underline hover:text-mauve-400">
            Register
          </Link>
        </p>
      </div>
    </div>
  )
}

export default LoginPage