/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,ts,tsx,js,jsx}'],
  theme: {
    extend: {
      colors: {
        stripe: {
          50: '#f0efff',
          100: '#e1deff',
          500: '#635bff',
          600: '#4f46e5',
          700: '#3e35d9',
        },
        // Stripe dark sidebar
        navy: {
          900: '#0a2540',
          800: '#0d2f50',
          700: '#1a3a5c',
        },
      },
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'system-ui', 'sans-serif'],
      },
      borderRadius: {
        '2xl': '1rem',
        '3xl': '1.5rem',
      },
      boxShadow: {
        'stripe-sm': '0 1px 3px rgba(60,66,87,0.08), 0 0 0 1px rgba(60,66,87,0.06)',
        'stripe-md': '0 4px 16px rgba(60,66,87,0.12), 0 0 0 1px rgba(60,66,87,0.06)',
        'stripe-lg': '0 20px 60px rgba(60,66,87,0.16)',
      },
    },
  },
  plugins: [],
};
