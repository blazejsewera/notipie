import React from 'react'

interface FootnoteProps {
  appName: string
  timestamp: string
}

const Footnote: React.FC<FootnoteProps> = ({ appName, timestamp }) => (
  <div>
    <p>{appName}</p>
    <p>{timestamp}</p>
  </div>
)

export default Footnote
