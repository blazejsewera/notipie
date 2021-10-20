import * as React from 'react'
import type { FC } from '../../type/react'
import { Moon, Sun } from '../../external/icon'
import { cx } from '../../util/cx'

export interface ToggleDarkModeButtonProps {
  isDarkMode: boolean
  onClick: () => void
}

export const ToggleDarkModeButton: FC<ToggleDarkModeButtonProps> = ({ isDarkMode, onClick }) => {
  const ToggleIcon = isDarkMode ? Sun : Moon
  return (
    <button onClick={onClick}>
      <ToggleIcon className={cx('stroke-current', 'text-gray-500', 'dark:text-gray-300', 'w-5', 'h-5')} />
    </button>
  )
}
