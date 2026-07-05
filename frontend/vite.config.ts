import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // Dev-only proxy to the Go backend; production sets VITE_API_BASE_URL.
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
})
