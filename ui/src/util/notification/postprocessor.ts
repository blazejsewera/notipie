import { Notification } from '../../type/notification'

export const addHandlers = (notification: Notification): Notification => ({
  ...notification,
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
})
