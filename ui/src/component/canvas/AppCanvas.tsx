import { useEffect } from 'react'
import { useStore, State } from '../../store'
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
          verticallyScrollable ? 'min-w-max' : '',
        )}
      >
        {children}
      </main>
      <AppControls isDarkMode={isDarkMode} onToggleDarkMode={toggleDarkMode} />
    </div>
  )
}

type SetPartialState = (partial: Partial<State>) => void
const checkForDarkModePreferenceConnected = (setState: SetPartialState) => {
  window.matchMedia('(prefers-color-scheme: dark)').matches
    ? setState({ darkMode: true })
    : setState({ darkMode: false })
}

type AppCanvasConnectedProps = Pick<AppCanvasProps, 'verticallyScrollable' | 'children'>
export const AppCanvasConnected: FC<AppCanvasConnectedProps> = ({ verticallyScrollable, children }) => {
  const { isDarkMode, toggleDarkMode } = useStore(state => ({
    isDarkMode: state.darkMode,
    toggleDarkMode: state.darkModeToggle,
  }))
  const checkForDarkModePreference = () => checkForDarkModePreferenceConnected(useStore.setState)
  return <AppCanvas {...{ isDarkMode, toggleDarkMode, verticallyScrollable, checkForDarkModePreference, children }} />
}
