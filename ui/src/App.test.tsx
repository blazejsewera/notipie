import React from 'react'
import { render } from '@testing-library/react'
import App from './App'

describe('App tests', () => {
  it('tests the app', () => {
    const { getByText } = render(<App />)
    const button = getByText('Hello')
    expect(button.textContent).toEqual('Hello')
  })
})
