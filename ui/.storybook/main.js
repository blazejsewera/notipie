module.exports = {
  framework: '@storybook/react',
  core: {
    builder: '@storybook/builder-vite',
  },
  features: {
    storyStoreV7: true,
  },
  staticDirs: ['../public'],
  stories: ['../story/**/*.story.mdx', '../story/**/*.@(js|jsx|ts|tsx)'],
  addons: ['@storybook/addon-links', '@storybook/addon-essentials', 'storybook-dark-mode', '@storybook/addon-postcss'],
}
