import express from 'express'
import expressWs from 'express-ws'
import { getMockNotifications } from './notification'
import { handleNotificationSending, initStdin } from './realtime'

const main = () => {
  const port = 8080
  const { app } = expressWs(express())

  app.get('/notifications', (_, res) => {
    res.send({ notifications: getMockNotifications(9) })
  })

  app.ws('/ws', ws => {
    handleNotificationSending(ws)
  })

  app.listen(port, () => {
    console.log(`Example server listening on port ${port}`)
  })
}

initStdin()
main()
