export type OnNotificationCardCheck = () => void
export type OnNotificationCardArchive = () => void
export type OnNotificationCardSettings = () => void

export interface NotificationCardHandlers {
  onCheck: OnNotificationCardCheck
  onArchive: OnNotificationCardArchive
  onSettings: OnNotificationCardSettings
}

export type OnNotificationContainerCheckAll = () => void

export interface NotificationContainerHandlers {
  onCheckAll: OnNotificationContainerCheckAll
}

export type NotificationContainerHandlersFactory = (category: string) => NotificationContainerHandlers
