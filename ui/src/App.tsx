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
import { getNotifications, ping } from './net/sync/api'

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

    getNotifications().then(n => console.log(n))
    ping().then(ok => console.log(ok))
  }, [])

  return (
    <AppCanvas verticallyScrollable>
      <NotificationBoard intl={intl} containerHandlersFactory={mockContainerHandlersFactory} />
    </AppCanvas>
  )
}
