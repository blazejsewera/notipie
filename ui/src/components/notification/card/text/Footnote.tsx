import * as React from 'react'

export interface FootnoteProps {
  appName: string
  timestamp: string
}

export const Footnote: React.FC<FootnoteProps> = ({ appName, timestamp }) => (
  <div>
    <p>{appName}</p>
    <p>{timestamp}</p>
  </div>
)
