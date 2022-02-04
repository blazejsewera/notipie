import { DarkmodeOff, DarkmodeOn, DarkmodeToggle } from './darkmode/set'
import {
  Request as NotificationFetchRequest,
  Success as NotificationFetchSuccess,
  Fail as NotificationFetchFail,
} from './notification/fetch'
import {
  ReceivePushed as NotificationReceivePushed,
  ReceivePushedError as NotificationReceivePushedError,
} from './notification/push'
import { UpdateTime as NotificationUpdateTime } from './notification/time'

export type Action =
  | NotificationFetchRequest
  | NotificationFetchSuccess
  | NotificationFetchFail
  | NotificationReceivePushed
  | NotificationReceivePushedError
  | NotificationUpdateTime
  | DarkmodeOn
  | DarkmodeOff
  | DarkmodeToggle

export {
  actionRequest as actionFetchRequest,
  actionSuccess as actionFetchSuccess,
  actionFail as actionFetchFail,
} from './notification/fetch'

export { actionReceivePushed } from './notification/push'

export { actionDarkmodeOn, actionDarkmodeOff, actionDarkmodeToggle } from './darkmode/set'
