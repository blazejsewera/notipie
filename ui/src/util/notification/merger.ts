import { Notification } from '../../type/notification'

type DuplicateDetector = (notifications: Notification[], current: Notification) => boolean

const isDuplicateByIdInArray: DuplicateDetector = (notifications, current) => {
  if (!current.id) return false
  return notifications.map(m => m.id ?? '').includes(current.id)
}

type NotificationMerger = (notifications: Notification[]) => Notification[]

/**
 * Merges (deduplicates) notification array based on notification id.
 * If an id is not present in the notification, it is not deduplicated.
 *
 * @example
 * const notifications = [
 *   { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
 *   { appName: 'b', timestamp: 'b', title: 'b', id: '2' },
 *   { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
 * ]
 * const merged = merge(notifications)
 * // will merge them into a 2-element array containing 0th and 1st element of msgs
 * // id: '1' is deduplicated
 */
export const merge: NotificationMerger = notifications =>
  notifications.reduce(
    (merged: Notification[], current) => (isDuplicateByIdInArray(merged, current) ? merged : [...merged, current]),
    [],
  )
