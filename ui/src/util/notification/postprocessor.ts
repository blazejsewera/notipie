import { Notification, NotificationWithHandlers } from '../../type/notification'
import { updateTimeAll } from './time'

export const addHandlers = (notifications: Notification[]): NotificationWithHandlers[] =>
  notifications.map(notification => ({
    notification,
    handlers: {
      onArchive: () => {
        notification.archiveUri && open(notification.archiveUri)
      },
      onCheck: () => {
        notification.readUri && open(notification.readUri)
      },
      onExternal: () => {
        notification.extUri && open(notification.extUri)
      },
    },
  }))

export const postprocessAll = (notificationsWithHandlers: NotificationWithHandlers[]): NotificationWithHandlers[] =>
  updateTimeAll(notificationsWithHandlers)

export const addHandlersAndPostprocess = (notifications: Notification[]): NotificationWithHandlers[] =>
  postprocessAll(addHandlers(notifications))
