export const T_DARKMODE_ON = 'darkmode/on'
export type DarkmodeOn = {
  type: typeof T_DARKMODE_ON
}
export const actionDarkmodeOn = (): DarkmodeOn => ({ type: T_DARKMODE_ON })

export const T_DARKMODE_OFF = 'darkmode/off'
export type DarkmodeOff = {
  type: typeof T_DARKMODE_OFF
}
export const actionDarkmodeOff = (): DarkmodeOff => ({ type: T_DARKMODE_OFF })

export const T_DARKMODE_TOGGLE = 'darkmode/toggle'
export type DarkmodeToggle = {
  type: typeof T_DARKMODE_TOGGLE
}
export const actionDarkmodeToggle = (): DarkmodeToggle => ({ type: T_DARKMODE_TOGGLE })
