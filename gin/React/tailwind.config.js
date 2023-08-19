/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  spacing: {
    128: '32rem',
    144: '36rem',
  },
  borderRadius: {
    '4xl': '2rem',
  },
  theme: {
    fontFamily: {
      sans: ['Oxygen-Regular', 'sans-serif'],
    },
    extend: {
      colors: {
        primary: '#434343',
        secondary: '#f7c873',
      },
    },
  },
  plugins: [],
};
