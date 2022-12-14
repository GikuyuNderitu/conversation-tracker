/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: ({ colors }) => ({
      inherit: colors.inherit,
      current: colors.current,
      transparent: colors.transparent,
      black: colors.black,
      white: colors.white,
      slate: colors.slate,
      gray: colors.gray,
      zinc: colors.zinc,
      neutral: colors.neutral,
      stone: colors.stone,
      red: colors.red,
      orange: colors.orange,
      amber: colors.amber,
      yellow: colors.yellow,
      lime: colors.lime,
      green: colors.green,
      emerald: colors.emerald,
      teal: colors.teal,
      cyan: colors.cyan,
      sky: colors.sky,
      blue: colors.blue,
      indigo: colors.indigo,
      violet: colors.violet,
      purple: colors.purple,
      fuchsia: colors.fuchsia,
      pink: colors.pink,
      rose: colors.rose,
      'primary': 'var(--primary-color)',
      'primary-surface': 'var(--primary-surface)',
      'secondary': 'var(--secondary-color)',
      'accent': 'var(--accent-color)',
      'on-surface': 'var(--on-surface)',
      'on-surface-text': 'var(--on-surface-text)',
      'on-surface-btn': 'var(--on-surface-btn)',
      'on-surface-secondary-btn': 'var(--on-surface-secondary-btn)',
      'on-surface-btn-text': 'var(--on-surface-btn-text)',
      'on-surface-secondary-btn-text': 'var(--on-surface-secondary-btn-text)',
    }),
    // Keep sorted.
    extend: {
      minHeight: {
        'tap-target': '48px',
      },
      minWidth: {
        'tap-target': '48px',
      },
      transitionProperty: {
        'border-radius': 'border-radius'
      },
    },
  },
  darkMode: {},
  plugins: [
    require('@tailwindcss/forms'),
    require('tailwind-scrollbar-hide'),
  ],
}
