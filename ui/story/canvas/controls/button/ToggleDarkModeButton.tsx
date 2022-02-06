import { Meta, Story } from '@storybook/react'
import {
  ToggleDarkModeButton,
  ToggleDarkModeButtonProps,
} from '../../../../src/component/canvas/controls/button/ToggleDarkModeButton'

export default {
  title: 'canvas/controls/button/ToggleDarkModeButton',
  component: ToggleDarkModeButton,
} as Meta

export const ExampleToggleDarkModeButton: Story<ToggleDarkModeButtonProps> = args => <ToggleDarkModeButton {...args} />

ExampleToggleDarkModeButton.args = {
  isDarkMode: true,
  onClick: () => {},
}
