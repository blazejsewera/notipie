import { cx } from './cx'

describe('tests for class names merge function', () => {
  it('returns empty string on no args', () => {
    const c = cx()
    expect(c).toEqual('')
  })

  it('returns one class name on one arg', () => {
    const c = cx('a')
    expect(c).toEqual('a')
  })

  it('returns space-delimited class names on more than one arg', () => {
    const c = cx('a', 'b', 'c')
    expect(c).toEqual('a b c')
  })
})
