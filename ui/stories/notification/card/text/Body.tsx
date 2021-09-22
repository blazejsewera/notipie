import React from 'react'
import { Story, Meta } from '@storybook/react'

import Body from '../../../../src/components/notification/card/text/Body'
import { fiveSentenceLoremIpsum } from '../../../assets/lipsum'

export default {
  title: 'notification/card/text/Body',
  component: Body,
} as Meta

export const ExampleBody: Story = (args) => <Body {...args} />

ExampleBody.args = {
  children: fiveSentenceLoremIpsum,
}
