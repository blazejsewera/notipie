import express from 'express'
import { getMockNotifications } from './notification'

const main = () => {
  const app = express()
  const port = 8080

  app.get('/notifications', (req, res) => {
    res.send({ notifications: getMockNotifications(9) })
  })

  app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
  })
}

main()
