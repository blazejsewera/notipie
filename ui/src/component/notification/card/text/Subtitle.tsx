import * as React from 'react'
import type { FCC } from '../../../../type/react'
import { cx } from '../../../../util/cx'

export type SubtitleChildren = string

export const Subtitle: FCC<SubtitleChildren> = ({ children }) =>
  children ? <p className={cx('text-xs', 'font-bold', 'text-gray-500')}>{children}</p> : null
