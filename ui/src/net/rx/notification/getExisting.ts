import { mergeMap, Observable, observableFrom } from '../../../external/rxjs'
import { Notification } from '../../../type/notification'
import { getNotifications } from '../../sync/api'

export const getExisting$: Observable<Notification> = observableFrom(getNotifications(fetch)).pipe(
  mergeMap(observableFrom),
)
