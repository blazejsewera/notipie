import { WebSocket, MessageEvent } from 'ws'

const main = () => {
  const logMessage = (event: MessageEvent) => {
    console.log('message received')
    console.log(`event type: ${event.type}`)
    console.log(`event data: ${event.data}`)
  }

  const protocol = 'ws'
  const address = 'localhost'
  const port = 8080
  const path = '/ws'

  const ws = new WebSocket(`${protocol}://${address}:${port}${path}`)

  ws.onerror = event => {
    console.error('error in websocket')
    console.error(event.message)
  }

  ws.onopen = () => {
    console.log('websocket conn open')
  }

  ws.onmessage = logMessage

  ws.onclose = () => {
    console.log('websocket conn closed')
  }
}

main()
