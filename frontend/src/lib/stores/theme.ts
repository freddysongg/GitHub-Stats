import { writable } from 'svelte/store';

function createThemeStore() {
  const { subscribe, set } = writable('light');

  return {
    subscribe,
    toggleTheme: () => {
      subscribe((theme) => {
        const newTheme = theme === 'light' ? 'dark' : 'light';
        document.documentElement.classList.toggle('dark');
        set(newTheme);
      });
    },
    initialize: () => {
      const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
      if (isDark) {
        document.documentElement.classList.add('dark');
        set('dark');
      }
    },
  };
}

export const theme = createThemeStore();
