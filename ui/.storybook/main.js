module.exports = {
  stories: ['../stories/**/*.stories.mdx', '../stories/**/*.@(js|jsx|ts|tsx)'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    {
      name: '@storybook/addon-postcss',
      options: {
        postcssLoaderOptions: {
          implementation: require('postcss'), // enable postcss@8 for tailwind
        },
      },
    },
  ],
}
