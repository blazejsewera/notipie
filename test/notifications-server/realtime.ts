import { stdin } from 'process'
import { emitKeypressEvents, Key } from 'readline'
import { WebSocket } from 'ws'
import { getMockNotification, Notification } from './notification'

export const initStdin = () => {
  if (stdin.isTTY) stdin.setRawMode(true)
  emitKeypressEvents(stdin)
}

const toJson = (notification: Notification): string => JSON.stringify(notification)

export const handleNotificationSending = (ws: WebSocket) => {
  console.log('> ws conn open, press enter to send a notification')

  const keypressHandler = (_: string, key: Key) => {
    if (key.name !== 'enter') return
    const notification = getMockNotification()
    ws.send(toJson(notification))
    console.log(`> notification sent, id = ${notification.id}`)
  }

  stdin.on('keypress', keypressHandler)

  ws.on('close', () => {
    console.log('> ws conn closed')
    stdin.removeListener('keypress', keypressHandler)
  })
}
