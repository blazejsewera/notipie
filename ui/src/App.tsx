import { useEffect } from 'react'
import { FC } from './type/react'
import { intl } from './i18l/intl'
import { handlersFactory as mockContainerHandlersFactory } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { AppCanvasConnected as AppCanvas } from './component/canvas/AppCanvas'
import { NotificationBoardConnected as NotificationBoard } from './component/notification/board/NotificationBoard'
import { useStore } from './store'
import { getNotifications } from './net/sync/api'
import { addHandlersAndPostprocess } from './util/notification/postprocessor'

export const App: FC = () => {
  const setNotifications = useStore(state => state.notificationFetchSuccess)
  useEffect(() => {
    getNotifications().then(n => setNotifications(addHandlersAndPostprocess(n)))
  }, [])

  return (
    <AppCanvas verticallyScrollable>
      <NotificationBoard intl={intl} containerHandlersFactory={mockContainerHandlersFactory} />
    </AppCanvas>
  )
}
