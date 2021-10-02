import * as React from 'react'
import { cx } from '../../../../util/cx'
import { Meta } from '../text/Meta'
import { Controls } from './Controls'

export interface FooterProps {
  appName: string
  timestamp: string
}

export const Footer: React.FC<FooterProps> = ({ appName, timestamp }) => (
  <div className={cx('flex')}>
    <Meta {...{ appName, timestamp }} />
    <Controls />
  </div>
)
