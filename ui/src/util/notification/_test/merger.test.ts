import { Notification } from '../../../type/notification'
import { merge } from '../merger'

describe('notification merger to reduce duplication', () => {
  it('deduplicates notifications', () => {
    // given
    const notifications: Notification[] = [
      { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
      { appName: 'b', timestamp: 'b', title: 'b', id: '2' },
      { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
      { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
    ]
    const merged: Notification[] = [
      { appName: 'a', timestamp: 'a', title: 'a', id: '1' },
      { appName: 'b', timestamp: 'b', title: 'b', id: '2' },
    ]

    // when
    const tested = merge(notifications)

    // then
    expect(tested.length).toEqual(merged.length)
    expect(tested).toEqual(merged)
  })
})
