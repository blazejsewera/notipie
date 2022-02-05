import { NotificationWithHandlers } from '../../../type/notification'

export const T_RECEIVE_PUSHED = 'notification/receivePushed'
export type ReceivePushed = {
  type: typeof T_RECEIVE_PUSHED
  notificationWithHandlers: NotificationWithHandlers
}
export const actionReceivePushed = (notificationWithHandlers: NotificationWithHandlers): ReceivePushed => ({
  type: T_RECEIVE_PUSHED,
  notificationWithHandlers,
})

export const T_RECEIVE_PUSHED_ERROR = 'notification/receivePushed/error'
export type ReceivePushedError = {
  type: typeof T_RECEIVE_PUSHED_ERROR
  message: string
}
export const actionReceivePushedError = (message: string): ReceivePushedError => ({
  type: T_RECEIVE_PUSHED_ERROR,
  message,
})
