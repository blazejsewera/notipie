import * as React from 'react'
import type { FC } from '../../../../type/react'
import { Archive, Check, Settings } from 'react-feather'
import { ControlIcon } from '../sprite/icon/ControlIcon'

export const Controls: FC = () => (
  <div className="inline-grid grid-cols-3 gap-x-5 ml-auto my-auto mr-2">
    <ControlIcon SvgIcon={Archive} />
    <ControlIcon SvgIcon={Check} />
    <ControlIcon SvgIcon={Settings} />
  </div>
)
