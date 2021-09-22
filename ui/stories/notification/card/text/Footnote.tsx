import React from 'react'
import { Story, Meta } from '@storybook/react'

import Footnote, { FootnoteProps } from '../../../../src/components/notification/card/text/Footnote'

export default {
  title: 'notification/card/text/Footnote',
  component: Footnote,
} as Meta

export const ExampleFootnote: Story<FootnoteProps> = (args) => <Footnote {...args} />

ExampleFootnote.args = {
  appName: 'Example App',
  timestamp: '2 hours ago',
}
