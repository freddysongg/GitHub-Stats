export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        github: {
          50: '#f6f8fa',
          100: '#ebeef1',
          200: '#d0d7de',
          300: '#8b949e',
          400: '#6e7681',
          500: '#484f58',
          600: '#30363d',
          700: '#21262d',
          800: '#161b22',
          900: '#0d1117',
        },
      },
      borderRadius: {
        mac: '0.75rem',
      },
      boxShadow: {
        mac: '0 2px 4px rgba(0, 0, 0, 0.1), 0 8px 16px rgba(0, 0, 0, 0.1)',
      },
    },
  },
  plugins: [],
};
