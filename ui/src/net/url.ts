import { config } from '../config'

const { address, port, prefix, root, push, webSocket, notifications } = config.endpointConfig
const host = `${address}:${port}`

const wsScheme = 'ws'
const httpScheme = 'http'
type Scheme = typeof wsScheme | typeof httpScheme

const getUrl = (scheme: Scheme, target: string): string => `${scheme}://${host}${prefix}${target}`

export const rootUrl = getUrl(httpScheme, root)
export const pushUrl = getUrl(httpScheme, push)
export const wsUrl = getUrl(wsScheme, webSocket)
export const notificationsUrl = getUrl(httpScheme, notifications)
