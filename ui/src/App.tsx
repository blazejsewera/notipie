import { FC } from './type/react'
import { intl } from './i18l/intl'
import { handlersFactory as mockContainerHandlersFactory } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { AppCanvasConnected as AppCanvas } from './component/canvas/AppCanvas'
import { NotificationBoardConnected as NotificationBoard } from './component/notification/board/NotificationBoard'

export const App: FC = () => (
  <AppCanvas verticallyScrollable>
    <NotificationBoard intl={intl} containerHandlersFactory={mockContainerHandlersFactory} />
  </AppCanvas>
)
