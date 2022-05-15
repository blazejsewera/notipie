import { Observable, webSocket } from '../../../external/rxjs'
import { Notification } from '../../../type/notification'
import { wsUrl } from '../../url'

export const receivePushed$: Observable<Notification> = webSocket(wsUrl)
