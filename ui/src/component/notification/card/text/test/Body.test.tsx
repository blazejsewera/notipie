import * as React from 'react'
import { fireEvent, render, screen } from '@testing-library/react'
import { Body } from '../Body'

describe('tests for Body component', () => {
  // given
  const text = 'abc'
  const textRegex = new RegExp(`^${text}`)

  it('should display show more button', () => {
    // when
    render(<Body trimThreshold={2}>{text}</Body>)

    // then
    const seeMoreButton = screen.getByLabelText('show more')
    expect(seeMoreButton).toBeTruthy()
  })

  it('should show the whole string after clicking show more button', () => {
    // given
    render(<Body trimThreshold={2}>{text}</Body>)

    // when
    const seeMoreButton = screen.getByLabelText('show more')
    fireEvent.click(seeMoreButton)

    // then
    const fullText = screen.getByText(text)
    expect(fullText.innerHTML).toMatch(textRegex)
  })
})
