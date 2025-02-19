import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
//import vueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/

export default defineConfig({
  base: process.env.NODE_ENV==='production'?"/book_family/":"/",
  build: {
    outDir: 'dist',
  },
  assetsDir: 'assets',
  plugins: [
    vue(),
    
   // vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    hmr: {
      overlay: false,
    },
  },
})
