import { FC } from '../../../../type/react'
import { Title } from '../text/Title'
import { cx } from '../../../../util/classname/cx'
import { CheckAllButton } from '../sprite/button/CheckAllButton'
import { OnNotificationContainerCheckAll } from '../../../../type/handler'

export interface HeaderProps {
  title: string
  onCheckAll: OnNotificationContainerCheckAll
}

export const Header: FC<HeaderProps> = ({ title, onCheckAll }) => (
  <div className={cx('flex', 'mb-4', 'mx-2')}>
    <Title>{title}</Title>
    <CheckAllButton onClick={onCheckAll} />
  </div>
)
