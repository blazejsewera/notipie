import { Notification } from '../../../type/notification'
import { byUsername } from '../categorizer'

describe('notification categorizer', () => {
  // given
  const notifications: Notification[] = [
    { appName: 'a', timestamp: '2021-11-27T20:00:00.000Z', title: 'a' },
    { appName: 'b', timestamp: '2021-11-27T20:01:00.000Z', title: 'b' },
    { appName: 'a', timestamp: '2021-11-27T20:02:00.000Z', title: 'c' },
    { appName: 'd', timestamp: '2021-11-27T20:03:00.000Z', title: 'd' },
    { appName: 'd', timestamp: '2021-11-27T20:04:00.000Z', title: 'e' },
  ]

  const categorizedByUsername: Record<string, Notification[]> = {
    a: [
      { appName: 'a', timestamp: '2021-11-27T20:00:00.000Z', title: 'a' },
      { appName: 'a', timestamp: '2021-11-27T20:02:00.000Z', title: 'c' },
    ],
    b: [{ appName: 'b', timestamp: '2021-11-27T20:01:00.000Z', title: 'b' }],
    d: [
      { appName: 'd', timestamp: '2021-11-27T20:03:00.000Z', title: 'd' },
      { appName: 'd', timestamp: '2021-11-27T20:04:00.000Z', title: 'e' },
    ],
  }

  it('correctly categorizes by appName', () => {
    // when
    const categorized = byUsername(notifications)

    // then
    expect(categorized).toEqual(categorizedByUsername)
  })
})
