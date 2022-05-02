import * as c from '../../notipie.config.json'

export type Config = {
  prod: boolean
  proxyConfig: {
    httpScheme: 'http' | 'https'
    wsScheme: 'ws' | 'wss'
    address: string
    port: number
    prefix: string
  }
  endpointConfig: {
    address: string
    port: number
    prefix: string
    root: string
    push: string
    webSocket: string
    notifications: string
  }
}

export const config = c as Config
