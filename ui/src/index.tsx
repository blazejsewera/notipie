import * as React from 'react'
import { createRoot } from 'react-dom/client'
import { App } from './App'
import { wire as wireDependencies } from './wire'
;(() => {
  wireDependencies()

  const containerId = 'root'
  const container = document.getElementById(containerId)
  if (container === null) {
    console.error(`could not find ${containerId} element in DOM, check the HTML template`)
    return
  }

  const root = createRoot(container)

  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  )
})()
