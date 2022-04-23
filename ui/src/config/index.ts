import c from '../../notipie.config.json'

type Config = {
  prod: boolean
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
