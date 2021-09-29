import * as React from 'react'

export interface InitialsAppAvatarProps {
  appName: string
  bgColor?: string
}

export const InitialsAppAvatar: React.FC<InitialsAppAvatarProps> = ({ appName, bgColor }) => {
  const backgroundColor = bgColor ? bgColor : generateBackgroundColor(appName)
  const fontSize = 0.5 // svg is scaled for 1px, so it's roughly half the height
  return (
    <svg viewBox="0 0 1 1">
      <rect width={1} height={1} fill={backgroundColor}></rect>
      <text
        x={0.5}
        y={0.5}
        style={{ fontWeight: 'bold', fontFamily: 'Inter', fontSize: fontSize }}
        fill="#FFF"
        textAnchor="middle"
        dy={(fontSize * 0.356).toFixed(3)}
      >
        {makeInitials(appName)}
      </text>
    </svg>
  )
}

// TODO: implement a color generator
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const generateBackgroundColor = (_seed: string): string => '#000'

const makeInitials = (name: string) => name.substr(0, 2)
