import { useEffect } from 'react'
import { connect } from 'react-redux'
import { dispatch } from '../../store/store'
import { actionDarkmodeOn, actionDarkmodeOff, actionDarkmodeToggle } from '../../store/action/action'
import { State } from '../../store/store'
import { FC } from '../../type/react'
import { cx } from '../../util/classname/cx'
import { AppControls } from './controls/AppControls'

export interface AppCanvasProps {
  verticallyScrollable?: boolean
  checkForDarkModePreference: () => void
  isDarkMode: boolean
  toggleDarkMode: () => void
  children?: React.ReactNode
}

/**
 * Default application canvas with dark mode switch button
 */
export const AppCanvas: FC<AppCanvasProps> = ({
  verticallyScrollable,
  checkForDarkModePreference,
  isDarkMode,
  toggleDarkMode,
  children,
}) => {
  useEffect(() => {
    checkForDarkModePreference()
  }, [])

  return (
    <div className={cx('App', 'min-h-screen', isDarkMode ? 'dark' : '')}>
      <main
        className={cx(
          'bg-gray-200',
          'dark:bg-gray-600',
          'min-h-screen',
          'py-8',
          'px-5',
          'sm:py-12',
          'sm:px-8',
          verticallyScrollable ? '' : 'min-w-max',
        )}
      >
        {children}
      </main>
      <AppControls isDarkMode={isDarkMode} onToggleDarkMode={toggleDarkMode} />
    </div>
  )
}

const toggleDarkModeConnected = () => {
  dispatch(actionDarkmodeToggle())
}

const checkForDarkModePreferenceConnected = () => {
  window.matchMedia('(prefers-color-scheme: dark)').matches
    ? dispatch(actionDarkmodeOn())
    : dispatch(actionDarkmodeOff())
}

type StateMapper = (
  state: State,
) => Pick<AppCanvasProps, 'isDarkMode' | 'toggleDarkMode' | 'checkForDarkModePreference'>
const mapState: StateMapper = state => ({
  isDarkMode: state.isDarkMode,
  toggleDarkMode: toggleDarkModeConnected,
  checkForDarkModePreference: checkForDarkModePreferenceConnected,
})

export const AppCanvasConnected = connect(mapState)(AppCanvas)
