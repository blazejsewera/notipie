import * as React from 'react'
import { cx } from '../../../../utils/cx'

export const Subtitle: React.FC = ({ children }) => (
  <p className={cx('text-xs', 'font-bold', 'text-gray-500')}>{children}</p>
)
