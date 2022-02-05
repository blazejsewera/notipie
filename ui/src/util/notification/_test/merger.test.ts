import { NotificationWithHandlers } from '../../../type/notification'
import { merge } from '../merger'
import { mockNotificationCardHandlers as mockHandlers } from '../../../mock/notification.mock'

describe('notification merger to reduce duplication', () => {
  it('deduplicates notifications', () => {
    // given
    const notifications: NotificationWithHandlers[] = [
      { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: mockHandlers },
      { notification: { appName: 'b', timestamp: 'b', title: 'b', id: '2' }, handlers: mockHandlers },
      { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: mockHandlers },
      { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: mockHandlers },
    ]
    const merged: NotificationWithHandlers[] = [
      { notification: { appName: 'a', timestamp: 'a', title: 'a', id: '1' }, handlers: mockHandlers },
      { notification: { appName: 'b', timestamp: 'b', title: 'b', id: '2' }, handlers: mockHandlers },
    ]

    // when
    const tested = merge(notifications)

    // then
    expect(tested.length).toEqual(merged.length)
    expect(tested).toEqual(merged)
  })
})
