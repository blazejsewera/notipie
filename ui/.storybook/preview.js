import '../src/style/main.css'
import '../src/style/inter.css'
export const parameters = {
  backgrounds: {
    default: 'light',
    values: [
      {
        name: 'light',
        value: '#f3f4f6',
      },
      {
        name: 'dark',
        value: '#1f2937',
      },
    ],
  },
  actions: { argTypesRegex: '^on[A-Z].*' },
}
