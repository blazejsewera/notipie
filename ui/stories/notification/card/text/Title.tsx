import React from 'react'
import { Story, Meta } from '@storybook/react'

import Title from '../../../../src/components/notification/card/text/Title'

export default {
  title: 'notification/card/text/Title',
  component: Title,
} as Meta

export const ExampleTitle: Story = (args) => <Title {...args} />

ExampleTitle.args = {
  children: 'Example Title',
}
