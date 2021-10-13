import * as React from 'react'
import type { FC } from './type/react'
import { intl } from './i18l/intl'
import { full, fullWithImage, minimal, partial } from './mock/notification.mock'
import './style/main.css'
import './style/inter.css'
import { NotificationContainer } from './component/notification/container/NotificationContainer'

export const App: FC = () => {
  const notifications = [full, fullWithImage, partial, minimal]
  return (
    <div className="App bg-gray-100 min-h-screen p-10">
      <NotificationContainer notifications={notifications} intl={intl} />
    </div>
  )
}
