import { FC } from './type/react'
import { intl } from './i18l/intl'
import {
  fullWithHandlers,
  fullWithImageWithHandlers,
  minimalWithHandlers,
  partialWithHandlers,
} from './mock/notification.mock'
import { handlers as mockContainerHandlers } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { NotificationContainer } from './component/notification/container/NotificationContainer'
import { AppCanvasConnected } from './component/canvas/AppCanvas'
import { Provider } from 'react-redux'
import { store } from './store/store'

export const App: FC = () => {
  const notificationsWithHandlers = [
    fullWithHandlers,
    fullWithImageWithHandlers,
    partialWithHandlers,
    minimalWithHandlers,
  ]
  return (
    <Provider store={store}>
      <AppCanvasConnected verticallyScrollable>
        <NotificationContainer
          title="Tag title"
          {...{ notificationsWithHandlers, handlers: mockContainerHandlers, intl }}
        />
      </AppCanvasConnected>
    </Provider>
  )
}
