import * as React from 'react'
import type { FC } from './type/react'
import { NotificationCard } from './component/notification/card/NotificationCard'
import './style/main.css'
import './style/inter.css'

export const App: FC = () => {
  const body = `#12 add some new amazing functionality

Closes #10. Changes both in 'core' and 'ui'. Needs additional work with this and that.`

  const notification = {
    appName: 'Github',
    title: 'New Pull Request',
    subtitle: '— notipie',
    body,
    timestamp: '2 hours ago',
  }

  return (
    <div className="App bg-gray-100 h-screen p-10">
      <NotificationCard notification={notification} />
    </div>
  )
}
