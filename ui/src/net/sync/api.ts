import { Notification } from '../../type/notification'
import { notificationsUrl, rootUrl } from '../url'

export type Response<T> = {
  status: number
  json: () => Promise<T>
}
export type Fetch<T = void> = (url: string) => Promise<Response<T>>

/**
 * Pings the backend to find out if it is running
 * @returns true if backend responds properly, false otherwise
 */
export const ping = async (fetch: Fetch): Promise<boolean> =>
  fetch(rootUrl)
    .then(res => res.status === 200)
    .catch(() => false)

export type NotificationsRes = { notifications: Notification[] }

export const getNotifications = async (fetch: Fetch<NotificationsRes>): Promise<Notification[]> =>
  fetch(notificationsUrl)
    .then(res => res.json())
    .then(data => data.notifications)
    .catch((err: Error) => {
      throw new Error(`get client notifications: ${err.message}`)
    })
