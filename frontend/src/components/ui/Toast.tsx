import { useEffect, useState } from 'react'

interface ToastProps {
  message: string
  onClose: () => void
}

function Toast({ message, onClose }: ToastProps) {
  const [isVisible, setIsVisible] = useState(false)

  useEffect(() => {
    if (!message) return

    setIsVisible(true)

    const hideTimer = setTimeout(() => setIsVisible(false), 2700)
    const closeTimer = setTimeout(() => onClose(), 3000)

    return () => {
      clearTimeout(hideTimer)
      clearTimeout(closeTimer)
    }
  }, [message])

  if (!message) return null

  return (
    <div
      className={`
        fixed top-6 left-1/2 -translate-x-1/2 z-50 rounded-lg bg-red-500 px-5 py-3 text-white shadow-xl
        transition-all duration-200
        ${isVisible ? 'opacity-100 translate-y-0' : 'opacity-0 -translate-y-3'}
      `}
    >
      <p>{message}</p>
    </div>
  )
}

export default Toast