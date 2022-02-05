import { NotificationWithHandlers } from '../../type/notification'

type NotificationCategorizer = (
  notificationsWithHandlers: NotificationWithHandlers[],
) => Record<string, NotificationWithHandlers[]>

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
export const byAppName: NotificationCategorizer = notificationsWithHandlers =>
  notificationsWithHandlers.reduce(
    (categorized: Record<string, NotificationWithHandlers[]>, current) => ({
      ...categorized,
      [current.notification.appName]: [...(categorized[current.notification.appName] ?? []), current],
    }),
    {},
  )
