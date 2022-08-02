import { StorybookViteConfig } from '@storybook/builder-vite'

const config: StorybookViteConfig = {
  framework: '@storybook/react',
  core: {
    builder: '@storybook/builder-vite',
    disableTelemetry: true,
  },
  async viteFinal(config, { configType }) {
    if (configType == 'PRODUCTION')
      return {
        ...config,
        build: {
          chunkSizeWarningLimit: 3000,
        },
        optimizeDeps: {
          include: ['storybook-dark-mode'],
        },
      }
    return {
      ...config,
      build: {
        chunkSizeWarningLimit: 5000,
      },
    }
  },
  features: {
    storyStoreV7: true,
  },
  staticDirs: ['../public'],
  stories: ['../story/**/*.story.mdx', '../story/**/*.@(js|jsx|ts|tsx)'],
  addons: ['@storybook/addon-links', '@storybook/addon-essentials', 'storybook-dark-mode', '@storybook/addon-postcss'],
}

export default config
