import React from 'react'

export interface InitialsAppAvatarProps {
  initials: string
}

const InitialsAppAvatar: React.FC<InitialsAppAvatarProps> = ({ initials }) => {
  const backgroundColor = '#000'
  const fontSize = 12
  return (
    <svg viewBox="0 0 1 1">
      <rect width={1} height={1} fill={backgroundColor}></rect>
      <text
        x={0.5}
        y={0.5}
        style={{ fontWeight: 'bold', fontFamily: 'Inter', fontSize: `${fontSize}px` }}
        fill="#FFF"
        textAnchor="middle"
        dy={(fontSize * 0.356).toFixed(3)}
      >
        {initials}
      </text>
    </svg>
  )
}

export default InitialsAppAvatar
