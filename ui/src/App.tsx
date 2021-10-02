import * as React from 'react'
import { NotificationCard } from './components/notification/card/NotificationCard'
import './styles/main.css'
import './styles/inter.css'

export const App: React.FC = () => {
  const body = `#12 add some new amazing functionality

Closes #10. Changes both in 'core' and 'ui'. Needs additional work with this and that.`

  return (
    <div className="App bg-gray-100 h-screen p-10">
      <NotificationCard
        appName="Github"
        title="New Pull Request"
        subtitle="— notipie"
        body={body}
        timestamp="2 hours ago"
      />
    </div>
  )
}
