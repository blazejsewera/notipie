import { notificationsUrl, pushUrl, rootUrl, wsUrl } from '../url'

jest.mock('../../config', () => ({
  config: {
    proxyConfig: {
      httpScheme: 'http',
      wsScheme: 'ws',
      address: 'localhost',
      port: 8080,
      prefix: '/',
    },
    endpointConfig: {
      root: '',
      notifications: 'notifications',
      push: 'push',
      webSocket: 'ws',
    },
    prod: false,
  },
}))

describe('api urls are correct', () => {
  it('returns notifications url', () => {
    expect(notificationsUrl).toBe('http://localhost:8080/notifications')
  })
  it('returns push url', () => {
    expect(pushUrl).toBe('http://localhost:8080/push')
  })
  it('returns root url', () => {
    expect(rootUrl).toBe('http://localhost:8080/')
  })
  it('returns webSocket url', () => {
    expect(wsUrl).toBe('ws://localhost:8080/ws')
  })
})
