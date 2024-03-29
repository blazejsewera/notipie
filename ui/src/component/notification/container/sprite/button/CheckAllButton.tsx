import { FC } from '../../../../../type/react'
import { CheckSquare } from '../../../../../external/icon'
import { cx } from '../../../../../util/classname/cx'
import { OnNotificationContainerCheckAll } from '../../../../../type/handler'

export interface CheckAllButtonProps {
  onClick: OnNotificationContainerCheckAll
}

export const CheckAllButton: FC<CheckAllButtonProps> = ({ onClick }) => (
  <button type="button" className={cx('inline ml-auto my-auto shrink-0')} onClick={onClick}>
    <CheckSquare className={cx('stroke-current', 'text-gray-500', 'dark:text-gray-300', 'w-5', 'h-5')} />
  </button>
)
