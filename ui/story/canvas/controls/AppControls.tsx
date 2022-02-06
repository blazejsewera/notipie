import { Meta, Story } from '@storybook/react'
import { AppControls, AppControlsProps } from '../../../src/component/canvas/controls/AppControls'

export default {
  title: 'canvas/controls/AppControls',
  component: AppControls,
} as Meta

export const ExampleAppControls: Story<AppControlsProps> = args => <AppControls {...args} />

ExampleAppControls.args = {
  isDarkMode: true,
  onToggleDarkMode: () => {},
}
