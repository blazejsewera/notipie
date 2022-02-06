import { useEffect } from 'react'
import { FC } from './type/react'
import { intl } from './i18l/intl'
import { handlersFactory as mockContainerHandlersFactory } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { AppCanvasConnected as AppCanvas } from './component/canvas/AppCanvas'
import { Provider } from 'react-redux'
import { NotificationBoardConnected as NotificationBoard } from './component/notification/board/NotificationBoard'
import { dispatch, store } from './store/store'
import { actionFetchSuccess } from './store/action/action'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  minimalWithHandlers,
  otherAppWithHandlers,
  partialWithHandlers,
} from './mock/notification.mock'

export const App: FC = () => {
  useEffect(() => {
    const notificationsWithHandlers = [
      fullWithHandlers,
      fullWithImageWithHandlers,
      partialWithHandlers,
      minimalWithHandlers,
      otherAppWithHandlers,
    ]

    dispatch(actionFetchSuccess(notificationsWithHandlers))
  }, [])

  return (
    <Provider store={store}>
      <AppCanvas verticallyScrollable>
        <NotificationBoard intl={intl} containerHandlersFactory={mockContainerHandlersFactory} />
      </AppCanvas>
    </Provider>
  )
}
