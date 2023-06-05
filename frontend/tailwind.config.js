/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/lib/**/*.{svelte,js,ts}', './src/routes/**/*.{svelte,js,ts}'],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}

