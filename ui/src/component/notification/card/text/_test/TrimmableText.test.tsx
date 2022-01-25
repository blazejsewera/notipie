import { fireEvent, render, screen } from '@testing-library/react'
import { TrimmableText as Trim } from '../TrimmableText'
import { intlMock } from '../../../../../mock/intl.mock'

describe('tests for TrimmableText component', () => {
  // given
  const text = 'abc'
  const textRegex = new RegExp(`^${text}`)
  const threshold = 2
  const intl = intlMock

  it('should display show more button', () => {
    // when
    render(<Trim {...{ threshold, intl }}>{text}</Trim>)

    // then
    const seeMoreButton = screen.getByLabelText('show more')
    expect(seeMoreButton).toBeTruthy()
  })

  it('should show the whole string after clicking show more button', () => {
    // given
    render(<Trim {...{ threshold, intl }}>{text}</Trim>)

    // when
    const seeMoreButton = screen.getByLabelText('show more')
    fireEvent.click(seeMoreButton)

    // then
    const fullText = screen.getByText(text)
    expect(fullText.innerHTML).toMatch(textRegex)
  })
})
