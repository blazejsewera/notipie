import { Notification } from '../../type/notification'
import { parseTimestamp } from './time'

export const sortByNewest = (unsorted: Notification[]): Notification[] =>
  unsorted.slice().sort((a, b) => {
    const aDate = parseTimestamp(a)
    const bDate = parseTimestamp(b)
    return bDate > aDate ? 1 : bDate < aDate ? -1 : 0
  })
