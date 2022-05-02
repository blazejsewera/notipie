import { Notification } from '../../../type/notification'
import { ping, getNotifications } from '../api'

describe('api interfaces', () => {
  afterEach(() => {
    // prettier-ignore
    (global.fetch as jest.Mock).mockClear()
  })

  describe('ping', () => {
    const mockPingFetch = (response: Promise<any>) => {
      global.fetch = jest.fn().mockImplementationOnce(() => response)
    }

    const statusOk = Promise.resolve({ status: 200 })
    const statusInternalServerError = Promise.resolve({ status: 500 })
    const statusNetworkError = Promise.reject()

    it('pings the backend successfully', async () => {
      // given
      mockPingFetch(statusOk)

      // when
      const ok = await ping()

      // then
      expect(ok).toBe(true)
    })

    it('fails to ping the backend due to an invalid response code', async () => {
      // given
      mockPingFetch(statusInternalServerError)

      // when
      const ok = await ping()

      // then
      expect(ok).toBe(false)
    })

    it('fails to ping the backend due to network error', async () => {
      // given
      mockPingFetch(statusNetworkError)

      // when
      const ok = await ping()

      // then
      expect(ok).toBe(false)
    })
  })

  describe('notifications', () => {
    const mockNotificationsResponse = (notifications: Notification[]) => {
      global.fetch = jest.fn().mockImplementation(() =>
        Promise.resolve({
          status: 200,
          json: () => Promise.resolve({ notifications }),
        }),
      )
    }

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

    it('gets empty notification list', async () => {
      // given
      mockNotificationsResponse([])

      // when
      const notifications = await getNotifications()

      // then
      expect(notifications).toStrictEqual([])
    })

    it('gets non-empty notification list', async () => {
      // given
      mockNotificationsResponse(testNotifications)

      // when
      const notifications = await getNotifications()

      // then
      expect(notifications).toStrictEqual(testNotifications)
    })
  })
})
