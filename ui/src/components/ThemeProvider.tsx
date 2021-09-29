import 'inter-ui/inter-hinted-latin.css'
import * as React from 'react'
import style from './theme-provider.module.css'

interface ThemeProviderProps {}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  return <div className={style.common}>{children}</div>
}
