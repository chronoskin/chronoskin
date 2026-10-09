// Applies the chosen theme before the page is drawn, so it does not flash:
// that is why this file is loaded on its own, without defer.
try {
  const theme = localStorage.getItem('chronoskin.theme');
  if (theme === 'light' || theme === 'dark') document.documentElement.dataset.theme = theme;
} catch {
  // Storage is blocked: the system's theme stays.
}
