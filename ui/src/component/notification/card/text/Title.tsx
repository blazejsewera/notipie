import * as React from 'react'
import type { FCC } from '../../../../type/react'
import { cx } from '../../../../util/cx'

export const Title: FCC = ({ children }) => (
  <p className={cx('text-lg', 'font-bold', 'mb-0', 'leading-5', 'text-gray-800')}>{children}</p>
)
