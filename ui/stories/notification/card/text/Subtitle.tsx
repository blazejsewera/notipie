import React from 'react'
import { Story, Meta } from '@storybook/react'

import Subtitle from '../../../../src/components/notification/card/text/Subtitle'

export default {
  title: 'notification/card/text/Subtitle',
  component: Subtitle,
} as Meta

export const ExampleSubtitle: Story = (args) => <Subtitle {...args} />

ExampleSubtitle.args = {
  children: 'Example Subtitle',
}
