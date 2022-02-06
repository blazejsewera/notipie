import { NotificationContainerHandlers, NotificationContainerHandlersFactory } from '../type/handler'

export const handlers: NotificationContainerHandlers = {
  onCheckAll: () => {},
}

export const handlersFactory: NotificationContainerHandlersFactory = category => ({
  onCheckAll: () => {
    console.log(category)
  },
})
