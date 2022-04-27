import { createStore } from 'redux'
import { Store } from 'redux'
import { Action } from './action/action'
import { NotificationWithHandlers } from '../type/notification'
import { T_FAIL, T_REQUEST, T_SUCCESS } from './action/notification/fetch'
import { T_RECEIVE_PUSHED, T_RECEIVE_PUSHED_ERROR } from './action/notification/push'
import { merge } from '../util/notification/merger'
import { T_DARKMODE_OFF, T_DARKMODE_ON, T_DARKMODE_TOGGLE } from './action/darkmode/set'
import { T_UPDATE_TIME } from './action/notification/time'
import { updateTimeAll } from '../util/notification/time'
import { config } from '../config'

export type State = {
  state: 'loading' | 'ok' | 'fail'
  notificationsWithHandlers: NotificationWithHandlers[]
  isDarkMode: boolean
}

const defaultState: State = {
  state: 'ok',
  notificationsWithHandlers: [],
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
        notificationsWithHandlers: merge([
          ...previousState.notificationsWithHandlers,
          ...action.notificationsWithHandlers,
        ]),
      } // PERF: possible room for optimization
    case T_FAIL:
      console.warn(action.message)
      return { ...previousState, state: 'fail' }
    case T_RECEIVE_PUSHED:
      return {
        ...previousState,
        state: 'ok',
        notificationsWithHandlers: merge([...previousState.notificationsWithHandlers, action.notificationWithHandlers]),
      } // PERF: possible room for optimization
    case T_RECEIVE_PUSHED_ERROR:
      console.warn(action.message)
      return { ...previousState, state: 'fail' }
    case T_UPDATE_TIME:
      return { ...previousState, notificationsWithHandlers: updateTimeAll(previousState.notificationsWithHandlers) }
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

/* eslint-disable */
const composeEnhancers = !config.prod
  ? (window as any).__REDUX_DEVTOOLS_EXTENSION__ && (window as any).__REDUX_DEVTOOLS_EXTENSION__()
  : undefined
/* eslint-enable */

export const store: Store<State, Action> = createStore(reducer, composeEnhancers)
export const { dispatch, subscribe } = store

//
import create from 'zustand'

export type BearState = {
  state: 'loading' | 'ok' | 'fail'
  notificationsWithHandlers: NotificationWithHandlers[]
  darkMode: boolean
  darkModeOn: () => void
  darkModeOff: () => void
  darkModeToggle: () => void
}

/* eslint-disable */
const useStore = create<BearState>()(set => ({
  state: 'ok',
  notificationsWithHandlers: [],
  darkMode: false,
  darkModeOn: () => set({ darkMode: true }),
  darkModeOff: () => set({ darkMode: false }),
  darkModeToggle: () => set(prev => ({ darkMode: !prev.darkMode })),
}))
/* eslint-enable */
