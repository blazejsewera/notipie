import create from 'zustand'
import { devtools } from 'zustand/middleware'
import { config } from '../config'
import { NotificationWithHandlers } from '../type/notification'
import { merge } from '../util/notification/merger'
import { updateTimeAll } from '../util/notification/time'

export type State = {
  state: 'loading' | 'ok' | 'fail'
  errorMessage: string
  notificationsWithHandlers: NotificationWithHandlers[]
  darkMode: boolean
  darkModeOn: () => void
  darkModeOff: () => void
  darkModeToggle: () => void
  notificationFetchLoading: () => void
  notificationFetchFail: (errorMessage: string) => void
  notificationFetchSuccess: (fetched: NotificationWithHandlers[]) => void
  notificationPush: (pushed: NotificationWithHandlers) => void
  notificationPushFail: (errorMessage: string) => void
  notificationUpdateTime: () => void
}

export type SetState = (
  partial: State | Partial<State> | ((state: State) => State | Partial<State>),
  replace?: boolean | undefined,
) => void
const store = (set: SetState): State => ({
  state: 'ok',
  errorMessage: '',
  notificationsWithHandlers: [],
  darkMode: false,
  darkModeOn: () => set({ darkMode: true }),
  darkModeOff: () => set({ darkMode: false }),
  darkModeToggle: () => set(prev => ({ darkMode: !prev.darkMode })),
  notificationFetchLoading: () => set({ state: 'loading', errorMessage: '' }),
  notificationFetchFail: (errorMessage: string) => set({ state: 'fail', errorMessage }),
  notificationFetchSuccess: (fetched: NotificationWithHandlers[]) =>
    set(prev => ({
      state: 'ok',
      errorMessage: '',
      notificationsWithHandlers: merge([...prev.notificationsWithHandlers, ...fetched]),
    })),
  notificationPush: (pushed: NotificationWithHandlers) =>
    set(prev => ({
      state: 'ok',
      errorMessage: '',
      notificationsWithHandlers: merge([...prev.notificationsWithHandlers, pushed]),
    })),
  notificationPushFail: (errorMessage: string) => set({ state: 'fail', errorMessage }),
  notificationUpdateTime: () =>
    set(prev => ({ notificationsWithHandlers: updateTimeAll(prev.notificationsWithHandlers) })),
})

const storeWithMiddleware = config.prod ? store : devtools(store)

export const useStore = create<State>()(storeWithMiddleware)

export type UseStore = typeof useStore
export type GetState = typeof useStore.getState
export type Subscribe = typeof useStore.subscribe
export type Destroy = typeof useStore.destroy
