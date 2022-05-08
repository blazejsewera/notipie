import { Meta, Story } from '@storybook/react'

export default {
  title: 'Main',
} as Meta

const Main = () => <p>Welcome to Notipie Storybook</p>

export const Introduction: Story = () => <Main />
