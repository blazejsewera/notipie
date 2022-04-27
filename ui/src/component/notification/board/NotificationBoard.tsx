import { Intl } from '../../../i18l/intl'
import { useStore } from '../../../store'
import { NotificationContainerHandlersFactory } from '../../../type/handler'
import { NotificationWithHandlers } from '../../../type/notification'
import { FC } from '../../../type/react'
import { cx } from '../../../util/classname/cx'
import { byAppName } from '../../../util/notification/categorizer'
import { NotificationContainer } from '../container/NotificationContainer'

export interface NotificationBoardProps {
  notificationsWithHandlers: NotificationWithHandlers[]
  containerHandlersFactory: NotificationContainerHandlersFactory
  intl: Intl
}

export const NotificationBoard: FC<NotificationBoardProps> = ({
  notificationsWithHandlers,
  intl,
  containerHandlersFactory,
}) => {
  const categorized = byAppName(notificationsWithHandlers) // PERF: possible room for optimization
  const appNames = Object.keys(categorized)
  const containers = appNames.map(appName => (
    <NotificationContainer
      key={appName}
      title={appName}
      notificationsWithHandlers={categorized[appName]}
      intl={intl}
      handlers={containerHandlersFactory(appName)}
      style={cx('inline-block')}
    />
  ))

  return <div className={cx('inline-block space-x-4 whitespace-nowrap')}>{containers}</div>
}

type NotificationBoardConnectedProps = Pick<NotificationBoardProps, 'intl' | 'containerHandlersFactory'>
export const NotificationBoardConnected: FC<NotificationBoardConnectedProps> = ({ intl, containerHandlersFactory }) => {
  const notificationsWithHandlers = useStore(state => state.notificationsWithHandlers)
  return <NotificationBoard {...{ intl, containerHandlersFactory, notificationsWithHandlers }} />
}
