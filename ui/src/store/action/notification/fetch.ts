import { Notification } from '../../../type/notification'

export const T_REQUEST = 'notification/fetch/request'
export type Request = {
  type: typeof T_REQUEST
}
export const actionRequest = (): Request => ({ type: T_REQUEST })

export const T_SUCCESS = 'notification/fetch/success'
export type Success = {
  type: typeof T_SUCCESS
  notifications: Notification[]
}
export const actionSuccess = (notifications: Notification[]): Success => ({ type: T_SUCCESS, notifications })

export const T_FAIL = 'notification/fetch/fail'
export type Fail = {
  type: typeof T_FAIL
  notification: string
}
export const actionFail = (notification: string): Fail => ({ type: T_FAIL, notification })
