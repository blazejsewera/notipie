export type Notification = {
  appName: string
  timestamp: string
  appImgUri?: string
  title: string
  subtitle?: string
  body?: string
  id?: string
  read?: boolean
  extUri?: string
  readUri?: string
  archiveUri?: string
}
