import { useEffect } from 'react'
import { FC } from './type/react'
import { intl } from './i18l/intl'
import { handlersFactory as mockContainerHandlersFactory } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { AppCanvasConnected as AppCanvas } from './component/canvas/AppCanvas'
import { NotificationBoardConnected as NotificationBoard } from './component/notification/board/NotificationBoard'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  minimalWithHandlers,
  otherAppWithHandlers,
  partialWithHandlers,
  readWithHandlers,
} from './mock/notification.mock'
import { useStore } from './store'

export const App: FC = () => {
  const setNotifications = useStore(state => state.notificationFetchSuccess)
  useEffect(() => {
    const notificationsWithHandlers = [
      fullWithHandlers,
      fullWithImageWithHandlers,
      partialWithHandlers,
      minimalWithHandlers,
      otherAppWithHandlers,
      readWithHandlers,
    ]

    setNotifications(notificationsWithHandlers)
  }, [])

  return (
    <AppCanvas verticallyScrollable>
      <NotificationBoard intl={intl} containerHandlersFactory={mockContainerHandlersFactory} />
    </AppCanvas>
  )
}
