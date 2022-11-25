/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{vue,ts}'
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          light: '#7ff4b0',
          dark: '#49c17b'
        }
      }
    },
  },
  plugins: [],
}
