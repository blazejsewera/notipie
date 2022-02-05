import { NotificationWithHandlers } from '../../type/notification'

type DuplicateDetector = (
  notificationsWithHandlers: NotificationWithHandlers[],
  current: NotificationWithHandlers,
) => boolean

const isDuplicateByIdInArray: DuplicateDetector = (notificationsWithHandlers, current) => {
  if (!current.notification.id) return false
  return notificationsWithHandlers.map(m => m.notification.id ?? '').includes(current.notification.id)
}

type NotificationMerger = (notificationsWithHandlers: NotificationWithHandlers[]) => NotificationWithHandlers[]

/**
 * Merges (deduplicates) notificationWithHandlers array based on notification id.
 * If an id is not present in the notification, it is not deduplicated.
 *
 * @example
 * const notificationsWithHandlers = [
 *   { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: {...} },
 *   { notification: { appName: 'b', timestamp: 'b', title: 'b', id: '2' }, handlers: {...} },
 *   { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: {...} },
 * ]
 * const merged = merge(notificationsWithHandlers)
 * // will merge them into a 2-element array containing 0th and 1st element of msgs
 * // id: '1' is deduplicated
 */
export const merge: NotificationMerger = notificationsWithHandlers =>
  notificationsWithHandlers.reduce(
    (merged: NotificationWithHandlers[], current) =>
      isDuplicateByIdInArray(merged, current) ? merged : [...merged, current],
    [],
  )
