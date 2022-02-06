import { Meta, Story } from '@storybook/react'
import { AppCanvas, AppCanvasProps } from '../../src/component/canvas/AppCanvas'

export default {
  title: 'canvas/AppCanvas',
  component: AppCanvas,
} as Meta

export const ExampleAppCanvas: Story<AppCanvasProps> = args => <AppCanvas {...args} />

ExampleAppCanvas.args = {
  verticallyScrollable: false,
  checkForDarkModePreference: () => {},
  isDarkMode: true,
  toggleDarkMode: () => {},
}
