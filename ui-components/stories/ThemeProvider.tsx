import 'inter-ui/inter-hinted-latin.css'
import React from 'react'
import style from './theme-provider.module.css'

interface ThemeProviderProps {}

const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  return <div className={style.common}>{children}</div>
}

export default ThemeProvider
