import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export const config = defineConfig({
  plugins: [react()],
})
