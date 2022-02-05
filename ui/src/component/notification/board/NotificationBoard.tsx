import { connect } from 'react-redux'
import { Intl } from '../../../i18l/intl'
import { State } from '../../../store/store'
import { NotificationContainerHandlers } from '../../../type/handler'
import { NotificationWithHandlers } from '../../../type/notification'
import { FC } from '../../../type/react'
import { cx } from '../../../util/classname/cx'
import { byAppName } from '../../../util/notification/categorizer'
import { NotificationContainer } from '../container/NotificationContainer'

export interface NotificationBoardProps {
  notificationsWithHandlers: NotificationWithHandlers[]
  containerHandlers: NotificationContainerHandlers
  intl: Intl
}

export const NotificationBoard: FC<NotificationBoardProps> = ({
  notificationsWithHandlers,
  intl,
  containerHandlers,
}) => {
  const categorized = byAppName(notificationsWithHandlers) // PERF: possible room for optimization
  const appNames = Object.keys(categorized)
  const containers = appNames.map(appName => (
    <NotificationContainer
      key={appName}
      title={appName}
      notificationsWithHandlers={categorized[appName]}
      intl={intl}
      handlers={containerHandlers}
      style={cx('inline-block')}
    />
  ))

  return <div className={cx('inline-block space-x-4 whitespace-nowrap')}>{containers}</div>
}

type StateMapper = (state: State) => Pick<NotificationBoardProps, 'notificationsWithHandlers'>
const mapState: StateMapper = state => ({
  notificationsWithHandlers: state.notificationsWithHandlers,
})

export const NotificationBoardConnected = connect(mapState)(NotificationBoard)
