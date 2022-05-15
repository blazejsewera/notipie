import { interval, map, mergeMap, observableFrom } from '../../external/rxjs'
import { Status } from '../../store'
import { ping } from '../sync/api'

const toStatus = (ok: boolean): Status => (ok ? 'ok' : 'error')

type Milliseconds = number

export const pingBackend$ = (period: Milliseconds) =>
  interval(period).pipe(
    mergeMap(() => observableFrom(ping(fetch))),
    map(toStatus),
  )
