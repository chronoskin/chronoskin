// The light and dark switch in the header. The choice is kept in this
// browser; until one is made the site follows the system.
document.addEventListener('alpine:init', () => {
  Alpine.data('theme', () => ({
    toggle() {
      const root = document.documentElement;
      const dark = root.dataset.theme ? root.dataset.theme === 'dark' : matchMedia('(prefers-color-scheme: dark)').matches;
      root.dataset.theme = dark ? 'light' : 'dark';
      try {
        localStorage.setItem('chronoskin.theme', root.dataset.theme);
      } catch {
        // The theme holds until the page is left.
      }
    },
  }));
});
