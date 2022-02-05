import { NotificationWithHandlers } from '../../../type/notification'

export const T_REQUEST = 'notification/fetch/request'
export type Request = {
  type: typeof T_REQUEST
}
export const actionRequest = (): Request => ({ type: T_REQUEST })

export const T_SUCCESS = 'notification/fetch/success'
export type Success = {
  type: typeof T_SUCCESS
  notificationsWithHandlers: NotificationWithHandlers[]
}
export const actionSuccess = (notificationsWithHandlers: NotificationWithHandlers[]): Success => ({
  type: T_SUCCESS,
  notificationsWithHandlers,
})

export const T_FAIL = 'notification/fetch/fail'
export type Fail = {
  type: typeof T_FAIL
  message: string
}
export const actionFail = (message: string): Fail => ({ type: T_FAIL, message })
