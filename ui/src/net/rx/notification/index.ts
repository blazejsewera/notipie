import { catchError, map, merge, Observable, observableOf } from '../../../external/rxjs'
import { NotificationWithStatus } from '../../../store'
import { Notification } from '../../../type/notification'
import { addHandlers } from '../../../util/notification/postprocessor'
import { updateTime } from '../../../util/notification/time'
import { getExisting$ } from './getExisting'
import { receivePushed$ } from './receivePushed'

const toNotificationWithStatus = (notification: Notification): NotificationWithStatus => ({
  notification,
  status: 'ok',
})

export const notification$: Observable<NotificationWithStatus> = merge(getExisting$, receivePushed$).pipe(
  map(addHandlers),
  map(updateTime),
  map(toNotificationWithStatus),
  catchError(() => observableOf<NotificationWithStatus>({ status: 'error' })),
)
