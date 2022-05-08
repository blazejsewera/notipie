import { Meta, Story } from '@storybook/react'

export default {
  title: 'Main',
} as Meta

const Bg = ({ children }) => <div className="py-8 px-5">{children}</div>
const Header = ({ children }) => <h1 className="text-2xl font-bold text-gray-800 dark:text-gray-300">{children}</h1>
const Txt = ({ children }) => <p className="text-gray-800 dark:text-gray-300">{children}</p>

const Main = () => (
  <Bg>
    <Header>Welcome to Notipie Storybook</Header>
    <Txt>Look below for component examples</Txt>
  </Bg>
)

export const Introduction: Story = () => <Main />
