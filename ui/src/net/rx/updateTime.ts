import { interval, map } from '../../external/rxjs'

type Milliseconds = number

export const updateTime$ = (period: Milliseconds) => interval(period).pipe(map(() => {}))
