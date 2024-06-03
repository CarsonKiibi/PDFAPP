/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        main: {
          background: '#2E2E2E',
          text: '#E0E0E0',
          textsecond: '#B0B0B0',
          accent: '#FFA726',
          highlight: '#FFCC80',
          border: '#4D4D4D',
          card: '#3A3A3A',
        },
      },
    },
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {
    themes: [
      {
        theme: {
          "primary": "#FFA726",
          "secondary": "#FFCC80",
          "accent": "#3A3A3A",
          "neutral": "#2E2E2E",
          "base-100": "#3A3A3A", // Change the background color here
          "info": "#2094f3",
          "success": "#009485",
          "warning": "#ff9900",
          "error": "#ff5724",
        }
      }
    ]
  }
}

