import * as React from 'react'
import { Archive } from '../sprite/icon/Archive'
import { Check } from '../sprite/icon/Check'
import { Settings } from '../sprite/icon/Settings'

export const Controls: React.FC = () => (
  <div className="inline-grid grid-cols-3 gap-x-5 ml-auto my-auto mr-2">
    <Archive />
    <Check />
    <Settings />
  </div>
)
