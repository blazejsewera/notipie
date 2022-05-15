import { notification$ } from '../net/rx/notification'
import { pingBackend$ } from '../net/rx/pingBackend'
import { updateTime$ } from '../net/rx/updateTime'
import { useStore } from '../store'

export const wire = () => {
  const pingBackendPeriod = 10000
  const updateTimePeriod = 5000
  const { notificationReceived, notificationUpdateTime, statusSet } = useStore.getState()

  notification$.subscribe(notificationReceived)
  updateTime$(updateTimePeriod).subscribe(notificationUpdateTime)
  pingBackend$(pingBackendPeriod).subscribe(statusSet)
}
