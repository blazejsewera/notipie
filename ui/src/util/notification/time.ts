import { Notification, NotificationWithHandlers } from '../../type/notification'

type Period = 's' | 'm' | 'h' | 'D' | 'W' | 'M' | 'Y'
type Since = Record<Period, number>
const periodSequence: Period[] = ['s', 'm', 'h', 'D', 'W', 'M', 'Y']

export const formatSince = (since: Since): string => {
  let result = ''
  periodSequence
    .slice() // make a copy
    .reverse()
    .forEach(period => {
      if (since[period] <= 0) return
      result += `${since[period]}${period} `
    })
  if (result == '') return 'now'
  return `${result.slice(0, -1)} ago`
}

export const since = (past: Date, present: Date): Since => {
  const offset = Math.floor((present.getTime() - past.getTime()) / 1000)
  const lengths: Since = { s: 60, m: 60, h: 24, D: 7, W: 4.35, M: 12, Y: 10000 }
  const result: Since = { s: 0, m: 0, h: 0, D: 0, W: 0, M: 0, Y: 0 }

  let offsetLeft = offset

  periodSequence.forEach(period => {
    result[period] = offsetLeft % lengths[period]
    offsetLeft -= result[period]
    offsetLeft = Math.floor(offsetLeft / lengths[period])
  })

  return result
}

const getRelativeTime = (notification: Notification): string | undefined => {
  const now = new Date()
  const notificationTime = new Date(notification.timestamp)
  if (notificationTime.toString() === 'Invalid Date') return undefined
  const s = since(notificationTime, now)
  return formatSince(s)
}

export const updateTime = (notificationWithHandlers: NotificationWithHandlers): NotificationWithHandlers => ({
  ...notificationWithHandlers,
  notification: {
    ...notificationWithHandlers.notification,
    relativeTime: getRelativeTime(notificationWithHandlers.notification),
  },
})

export const updateTimeAll = (notifications: NotificationWithHandlers[]): NotificationWithHandlers[] =>
  notifications.map(updateTime)
