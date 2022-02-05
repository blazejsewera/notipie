import { NotificationWithHandlers } from '../../../type/notification'
import { byAppName } from '../categorizer'
import { mockNotificationCardHandlers as mockHandlers } from '../../../mock/notification.mock'

describe('notification categorizer', () => {
  // given
  const notificationsWithHandlers: NotificationWithHandlers[] = [
    { notification: { appName: 'a', timestamp: '2021-11-27T20:00:00.000Z', title: 'a' }, handlers: mockHandlers },
    { notification: { appName: 'b', timestamp: '2021-11-27T20:01:00.000Z', title: 'b' }, handlers: mockHandlers },
    { notification: { appName: 'a', timestamp: '2021-11-27T20:02:00.000Z', title: 'c' }, handlers: mockHandlers },
    { notification: { appName: 'd', timestamp: '2021-11-27T20:03:00.000Z', title: 'd' }, handlers: mockHandlers },
    { notification: { appName: 'd', timestamp: '2021-11-27T20:04:00.000Z', title: 'e' }, handlers: mockHandlers },
  ]

  const categorizedByUsername: Record<string, NotificationWithHandlers[]> = {
    a: [
      { notification: { appName: 'a', timestamp: '2021-11-27T20:00:00.000Z', title: 'a' }, handlers: mockHandlers },
      { notification: { appName: 'a', timestamp: '2021-11-27T20:02:00.000Z', title: 'c' }, handlers: mockHandlers },
    ],
    b: [{ notification: { appName: 'b', timestamp: '2021-11-27T20:01:00.000Z', title: 'b' }, handlers: mockHandlers }],
    d: [
      { notification: { appName: 'd', timestamp: '2021-11-27T20:03:00.000Z', title: 'd' }, handlers: mockHandlers },
      { notification: { appName: 'd', timestamp: '2021-11-27T20:04:00.000Z', title: 'e' }, handlers: mockHandlers },
    ],
  }

  it('correctly categorizes by appName', () => {
    // when
    const categorized = byAppName(notificationsWithHandlers)

    // then
    expect(categorized).toEqual(categorizedByUsername)
  })
})
