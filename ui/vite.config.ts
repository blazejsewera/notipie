import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import notipieConfig from './notipie.config.json'

const { address: endpointAddress, port: endpointPort } = notipieConfig.endpointConfig
const { httpScheme, wsScheme, address, port, prefix } = notipieConfig.proxyConfig
const rewrite = (path: string) => path.replace(new RegExp(`^${prefix}`), '/')

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port,
    host: address,
    proxy: {
      [`${prefix}ws`]: {
        target: `${wsScheme}://${endpointAddress}:${endpointPort}`,
        changeOrigin: true,
        secure: false,
        ws: true,
        rewrite,
      },
      [prefix]: {
        target: `${httpScheme}://${endpointAddress}:${endpointPort}`,
        changeOrigin: true,
        secure: false,
        rewrite,
      },
    },
  },
  preview: {
    host: '0.0.0.0',
    port: 5000,
  },
})
