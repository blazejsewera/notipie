import { Notification } from '../../type/notification'

type NotificationCategorizer = (notifications: Notification[]) => Record<string, Notification[]>

/**
 * Categorizes notifications by username, and puts them in an object, where
 * appName is the key, and notifications are the values.
 *
 * @example
 * const msgs = [
 *   { appName: 'a', timestamp: 'a', title: 'a' },
 *   { appName: 'a', timestamp: 'b', title: 'b' },
 *   { appName: 'b', timestamp: 'b', title: 'b' },
 * ]
 * const categorized = byUsername(msgs)
 * // will group them into { a: [...(2 msgs)], b: [...(1 msg)] }
 */
export const byUsername: NotificationCategorizer = notifications =>
  notifications.reduce(
    (categorized: Record<string, Notification[]>, current) => ({
      ...categorized,
      [current.appName]: [...(categorized[current.appName] ?? []), current],
    }),
    {},
  )
