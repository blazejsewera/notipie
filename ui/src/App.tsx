import * as React from 'react'
import type { FC } from './type/react'
import { intl } from './i18l/intl'
import { full, fullWithImage, minimal, partial } from './mock/notification.mock'
import { handlers as mockContainerHandlers } from './mock/notificationContainer.mock'
import './style/main.css'
import './style/inter.css'
import { NotificationContainer } from './component/notification/container/NotificationContainer'

export const App: FC = () => {
  const notifications = [full, fullWithImage, partial, minimal]
  return (
    <div className="App bg-gray-200 min-h-screen p-10">
      <NotificationContainer title="Tag title" {...{ notifications, handlers: mockContainerHandlers, intl }} />
    </div>
  )
}
