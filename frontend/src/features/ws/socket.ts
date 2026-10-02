import type { WSMessage } from './types'

type MessageHandler = (message: WSMessage) => void

export function createSocket(onMessage: MessageHandler) {
  const url = import.meta.env.VITE_WS_URL
  const socket = new WebSocket(url)

  socket.onmessage = (event) => {
    const message: WSMessage = JSON.parse(event.data)
    onMessage(message)
  }

  socket.onerror = (event) => {
    console.error('WebSocket error', event)
  }

  function sendClick() {
    if (socket.readyState !== WebSocket.OPEN) {
      console.warn('socket not ready yet, click ignored')
      return
    }
    socket.send(JSON.stringify({ action: 'click' }))
  }

  function close() {
    socket.close()
  }

  return { sendClick, close, socket }
}