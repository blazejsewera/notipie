import React from 'react'

export interface ButtonProps {
  size: 'small' | 'medium' | 'large'
  onClick: () => void
  label: string
  primary?: boolean
}

export const Button: React.FC<ButtonProps> = () => {
  return <button></button>
}
