// Keeps the scroll position from one style page to the next, so that on a
// phone the controls stay under the thumb when a step loads a new page.
(() => {
  const KEY = 'chronoskin.scroll';
  try {
    const kept = Number(sessionStorage.getItem(KEY));
    sessionStorage.removeItem(KEY);
    if (!location.pathname.startsWith('/s/')) return;
    if (kept > 0) scrollTo(0, kept);
    addEventListener('pagehide', () => sessionStorage.setItem(KEY, String(scrollY)));
  } catch {
    // Storage is blocked: pages open at their top.
  }
})();
