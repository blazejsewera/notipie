import { Notification } from '../../../type/notification'
import { sortByNewest } from '../sorter'

const sample: Notification = {
  appName: 'Sample AppName',
  title: 'Sample Title',
  timestamp: '',
}

const withTimestamp = (timestamp: string): Notification => ({ ...sample, timestamp })

describe('notification sorter', () => {
  it('sorts by newest', () => {
    // given
    const unsorted: Notification[] = [
      withTimestamp('2022-08-09T00:00:00.000Z'),
      withTimestamp('2022-08-08T23:00:00.000Z'),
      withTimestamp('2022-08-09T01:00:00.000Z'),
    ]
    const want: Notification[] = [
      withTimestamp('2022-08-09T01:00:00.000Z'),
      withTimestamp('2022-08-09T00:00:00.000Z'),
      withTimestamp('2022-08-08T23:00:00.000Z'),
    ]

    // when
    const have = sortByNewest(unsorted)

    // then
    expect(have).toStrictEqual(want)
  })
})
