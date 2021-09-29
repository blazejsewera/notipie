import * as React from 'react'
import { cx } from '../../../../utils/cx'

export const Title: React.FC = ({ children }) => (
  <p className={cx('text-lg', 'font-bold', 'mb-0', 'leading-5', 'text-gray-800')}>{children}</p>
)
