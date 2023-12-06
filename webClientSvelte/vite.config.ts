import { svelte } from '@sveltejs/vite-plugin-svelte'
import Icons from 'unplugin-icons/vite'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 8081
  },
  plugins: [
    svelte(),
    Icons({
      compiler: 'svelte',
      autoInstall: true,
    }),
  ],
})
