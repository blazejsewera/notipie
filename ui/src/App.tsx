import React, { useState, useEffect } from 'react'
import './main.css'
import { cx } from './utils/cx'
import { Button } from './components/Button'
import { AppAvatar } from './components/notification/card/avatar/AppAvatar'

export const App: React.FC = () => {
  // Create the count state.
  const [count, setCount] = useState(0)
  // Update the count (+1 every second).
  useEffect(() => {
    const timer = setTimeout(() => setCount(count + 1), 1000)
    return () => clearTimeout(timer)
  }, [count, setCount])
  // Return the App component.
  return (
    <div className="App">
      <header className={cx('App-header', 'flex')}>
        <p className={cx('text-xl', 'mx-auto')}>
          Page has been open for <code>{count}</code> seconds.
        </p>
      </header>
      <button>Hello</button>
      <Button size="small" onClick={() => null} label="">
        button
      </Button>
      <br />
      <br />
      <AppAvatar appName="Testapp" size="small" bgColor="red" />
      <br />
      <br />
      <AppAvatar appName="Testapp" size="medium" />
      <br />
      <br />
      <AppAvatar appName="Testapp" size="large" />
    </div>
  )
}
