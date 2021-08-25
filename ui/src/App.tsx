import React, { useState, useEffect } from 'react'
import './index.css'
import cx from './components/utils/cx'

const App: React.FC = () => {
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
    </div>
  )
}

export default App
