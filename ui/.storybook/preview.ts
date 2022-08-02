import '../src/style/main.css'
import '../src/style/inter.css'

export const parameters = {
  backgrounds: {
    values: [
      {
        name: 'light',
        value: '#E5E7EB',
      },
      {
        name: 'dark',
        value: '#4B5563',
      },
    ],
  },
  actions: { argTypesRegex: '^on[A-Z].*' },
  darkMode: {
    current: 'dark',
    stylePreview: true,
  },
}
