import { NotificationCardHandlers } from './handler'

export type Notification = {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
  id?: string
  handlers: NotificationCardHandlers
}

export type NotificationWithHandlers = {
  notification: Notification
  handlers: NotificationCardHandlers
}
