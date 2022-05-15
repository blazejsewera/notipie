import { Notification } from '../../../type/notification'
import { ping, getNotifications, Fetch, Response, NotificationsRes } from '../api'

describe('api interfaces', () => {
  describe('ping', () => {
    const mockFetch =
      (response: Promise<Response<void>>): Fetch =>
      () =>
        response

    const statusOk = Promise.resolve({ status: 200, json: () => Promise.resolve() })
    const statusInternalServerError = Promise.resolve({ status: 500, json: () => Promise.resolve() })
    const statusNetworkError = Promise.reject()

    it('pings the backend successfully', async () => {
      // given
      const fetch = mockFetch(statusOk)

      // when
      const ok = await ping(fetch)

      // then
      expect(ok).toBe(true)
    })

    it('fails to ping the backend due to an invalid response code', async () => {
      // given
      const fetch = mockFetch(statusInternalServerError)

      // when
      const ok = await ping(fetch)

      // then
      expect(ok).toBe(false)
    })

    it('fails to ping the backend due to network error', async () => {
      // given
      const fetch = mockFetch(statusNetworkError)

      // when
      const ok = await ping(fetch)

      // then
      expect(ok).toBe(false)
    })
  })

  describe('notifications', () => {
    const testNotifications: Notification[] = [
      {
        appName: 'TestAppName',
        timestamp: '2022-05-02T21:55:00.000Z',
        title: 'Test Title 1',
      },
      {
        appName: 'TestAppName',
        timestamp: '2022-05-02T21:56:00.000Z',
        title: 'Test Title 2',
      },
    ]

    const mockFetch =
      (notifications: Notification[]): Fetch<NotificationsRes> =>
      () =>
        Promise.resolve({ status: 200, json: () => Promise.resolve({ notifications }) })

    it('gets empty notification list', async () => {
      // given
      const fetch = mockFetch([])

      // when
      const notifications = await getNotifications(fetch)

      // then
      expect(notifications).toStrictEqual([])
    })

    it('gets non-empty notification list', async () => {
      // given
      const fetch = mockFetch(testNotifications)

      // when
      const notifications = await getNotifications(fetch)

      // then
      expect(notifications).toStrictEqual(testNotifications)
    })
  })
})
