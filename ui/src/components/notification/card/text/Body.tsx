import * as React from 'react'
import { cx } from '../../../../utils/cx'

export const Body: React.FC = ({ children }) => (
  <p className={cx('whitespace-pre-line', 'text-sm', 'leading-4', 'my-3', 'text-gray-800')}>{children}</p>
)
