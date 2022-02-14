export type OnNotificationCardCheck = () => void
export type OnNotificationCardArchive = () => void
export type OnNotificationCardExternal = () => void

export interface NotificationCardHandlers {
  onCheck: OnNotificationCardCheck
  onArchive: OnNotificationCardArchive
  onExternal: OnNotificationCardExternal
}

export type OnNotificationContainerCheckAll = () => void

export interface NotificationContainerHandlers {
  onCheckAll: OnNotificationContainerCheckAll
}

export type NotificationContainerHandlersFactory = (category: string) => NotificationContainerHandlers
