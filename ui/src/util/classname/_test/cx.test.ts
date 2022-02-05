import { cx } from '../cx'

describe('class names merge function', () => {
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

  it('returns space-delimited class names for mixed list and non list args', () => {
    const c = cx(['a'], ['b', 'c', ['d', 'e']], 'f')
    expect(c).toEqual('a b c d e f')
  })
})
