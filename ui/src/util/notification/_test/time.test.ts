import { since, formatSince } from '../time'

describe('time formatter for relative timestamps', () => {
  it('correctly calculates since', () => {
    // given
    const past = new Date('2021-01-01T00:00:00.000Z')
    const present = new Date('2021-01-01T01:02:03.654Z')
    const expected = { s: 3, m: 2, h: 1, D: 0, W: 0, M: 0, Y: 0 }

    // when
    const actual = since(past, present)

    // then
    expect(actual).toEqual(expected)
  })

  it('correctly formats since', () => {
    // given
    const since = { s: 10, m: 2, h: 0, D: 0, W: 0, M: 0, Y: 0 }

    // when
    const formatted = formatSince(since)

    // then
    expect(formatted).toEqual('2m 10s ago')
  })

  it('correctly formats now', () => {
    // given
    const since = { s: 0, m: 0, h: 0, D: 0, W: 0, M: 0, Y: 0 }

    // when
    const formatted = formatSince(since)

    // then
    expect(formatted).toEqual('now')
  })
})
