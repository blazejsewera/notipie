import { Notification } from '../../../type/notification'

export const T_RECEIVE_PUSHED = 'notification/receivePushed'
export type ReceivePushed = {
  type: typeof T_RECEIVE_PUSHED
  notification: Notification
}
export const actionReceivePushed = (notification: Notification): ReceivePushed => ({
  type: T_RECEIVE_PUSHED,
  notification,
})

export const T_RECEIVE_PUSHED_ERROR = 'notification/receivePushed/error'
export type ReceivePushedError = {
  type: typeof T_RECEIVE_PUSHED_ERROR
  notification: string
}
export const actionReceivePushedError = (notification: string): ReceivePushedError => ({
  type: T_RECEIVE_PUSHED_ERROR,
  notification,
})
