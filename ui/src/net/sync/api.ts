import { Notification } from '../../type/notification'
import { notificationsUrl, rootUrl } from '../url'

/**
 * Pings the backend to find out if it is running
 * @returns true if backend responds properly, false otherwise
 */
export const ping = async (): Promise<boolean> =>
  fetch(rootUrl)
    .then(res => res.status === 200)
    .catch(() => false)

type NotificationsRes = { notifications: Notification[] }

export const getNotifications = async (): Promise<Notification[]> =>
  fetch(notificationsUrl)
    .then(res => res.json())
    .then((data: NotificationsRes) => data.notifications)
    .catch((err: Error) => {
      throw new Error(`get client notifications: ${err.message}`)
    })
