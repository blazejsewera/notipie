import create from 'zustand'
import { devtools } from 'zustand/middleware'
import { config } from '../config'
import { Notification } from '../type/notification'
import { merge } from '../util/notification/merger'
import { sortByNewest } from '../util/notification/sorter'
import { updateTimeAll } from '../util/notification/time'

export type Status = 'ok' | 'error'
export type NotificationWithStatus = { notification?: Notification; status: Status }

export type State = {
  status: Status
  errorMessage: string
  notifications: Notification[]
  darkMode: boolean
  statusSet: (status: Status) => void
  darkModeOn: () => void
  darkModeOff: () => void
  darkModeToggle: () => void
  notificationReceived: (n: NotificationWithStatus) => void
  notificationUpdateTime: () => void
}

export type SetState = (
  partial: State | Partial<State> | ((state: State) => State | Partial<State>),
  replace?: boolean | undefined,
) => void
const store = (set: SetState): State => ({
  status: 'ok',
  errorMessage: '',
  notifications: [],
  darkMode: false,
  statusSet: backendState => set({ status: backendState }),
  darkModeOn: () => set({ darkMode: true }),
  darkModeOff: () => set({ darkMode: false }),
  darkModeToggle: () => set(prev => ({ darkMode: !prev.darkMode })),
  notificationReceived: n =>
    set(prev =>
      n.notification
        ? { status: n.status, notifications: sortByNewest(merge([n.notification, ...prev.notifications])) }
        : { status: n.status },
    ),
  notificationUpdateTime: () => set(prev => ({ notifications: updateTimeAll(prev.notifications) })),
})

const storeWithMiddleware = config.prod ? store : devtools(store)

export const useStore = create<State>()(storeWithMiddleware)

export type UseStore = typeof useStore
export type GetState = typeof useStore.getState
export type Subscribe = typeof useStore.subscribe
export type Destroy = typeof useStore.destroy
