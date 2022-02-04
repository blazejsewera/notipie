import { createStore } from 'redux'
import { Store } from 'redux'
import { Action } from './action/action'
import { Notification } from '../type/notification'
import { T_FAIL, T_REQUEST, T_SUCCESS } from './action/notification/fetch'
import { T_RECEIVE_PUSHED, T_RECEIVE_PUSHED_ERROR } from './action/notification/push'
import { merge } from '../util/notification/merger'
import { T_DARKMODE_OFF, T_DARKMODE_ON, T_DARKMODE_TOGGLE } from './action/darkmode/set'
import { T_UPDATE_TIME } from './action/notification/time'
import { updateTimeAll } from '../util/notification/time'
import { config } from '../config/config'

export type State = {
  state: 'loading' | 'ok' | 'fail'
  notifications: Notification[]
  notificationForm: Notification
  isDarkMode: boolean
}

const emptyNotification: Notification = {
  title: '',
  subtitle: '',
  body: '',
  appName: '',
  timestamp: '',
}

const defaultState: State = {
  state: 'ok',
  notifications: [],
  notificationForm: emptyNotification,
  isDarkMode: false,
}

type Reducer = (state: State | undefined, action: Action) => State

const reducer: Reducer = (previousState = defaultState, action) => {
  switch (action.type) {
    case T_REQUEST:
      return { ...previousState, state: 'loading' }
    case T_SUCCESS:
      return {
        ...previousState,
        state: 'ok',
        notifications: merge([...previousState.notifications, ...action.notifications]),
      } // PERF: possible room for optimization
    case T_FAIL:
      console.warn(action.notification)
      return { ...previousState, state: 'fail' }
    case T_RECEIVE_PUSHED:
      return {
        ...previousState,
        state: 'ok',
        notifications: merge([...previousState.notifications, action.notification]),
      } // PERF: possible room for optimization
    case T_RECEIVE_PUSHED_ERROR:
      console.warn(action.notification)
      return { ...previousState, state: 'fail' }
    case T_UPDATE_TIME:
      return { ...previousState, notifications: updateTimeAll(previousState.notifications) }
    case T_DARKMODE_ON:
      return { ...previousState, isDarkMode: true }
    case T_DARKMODE_OFF:
      return { ...previousState, isDarkMode: false }
    case T_DARKMODE_TOGGLE:
      return { ...previousState, isDarkMode: !previousState.isDarkMode }
    default:
      return previousState
  }
}

const isDev = config.mode === 'dev'
/* eslint-disable */
const composeEnhancers = isDev
  ? (window as any).__REDUX_DEVTOOLS_EXTENSION__ && (window as any).__REDUX_DEVTOOLS_EXTENSION__()
  : undefined
/* eslint-enable */

export const store: Store<State, Action> = createStore(reducer, composeEnhancers)
export const { dispatch, subscribe } = store
