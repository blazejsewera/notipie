import React from 'react'
import '../main.css'

export interface ButtonProps {
  size: 'small' | 'medium' | 'large'
  onClick: () => void
  label: string
  primary?: boolean
}

export const Button: React.FC<ButtonProps> = (props) => {
  return <button className="mx-2 ring-gray-500 ring-offset-indigo-700">{props.children}</button>
}
