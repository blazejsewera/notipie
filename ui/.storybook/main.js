module.exports = {
  staticDirs: ['../public'],
  stories: ['../story/**/*.story.mdx', '../story/**/*.@(js|jsx|ts|tsx)'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    'storybook-dark-mode',
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
