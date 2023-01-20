import { defineConfig, splitVendorChunkPlugin } from 'vite'
import autoPreprocess from 'svelte-preprocess';
import { svelte } from '@sveltejs/vite-plugin-svelte';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [splitVendorChunkPlugin(), svelte({
    preprocess: autoPreprocess()
  })],
  build: {
    manifest: "manifest.json"
  }
})
