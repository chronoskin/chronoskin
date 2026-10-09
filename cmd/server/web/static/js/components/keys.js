// Keyboard shortcuts: a control marked data-key="r" is clicked when R is
// pressed. Keys are left alone while typing, with a modifier held, or with
// a dialog or a list open.
function press(key) {
  if (document.querySelector('dialog[open], details[open]')) return false;
  const control = document.querySelector('[data-key="' + CSS.escape(key) + '"]:not([aria-disabled="true"])');
  if (!control) return false;
  control.click();
  return true;
}

document.addEventListener('keydown', (event) => {
  if (event.ctrlKey || event.metaKey || event.altKey || event.defaultPrevented) return;
  if (event.target.closest('input, textarea, select, [contenteditable]')) return;
  const key = event.key.length === 1 ? event.key.toLowerCase() : event.key;
  if (press(key)) event.preventDefault();
});

// After a click into the preview, keys go to the specimen; it hands the
// shortcut keys back up.
window.addEventListener('message', (event) => {
  const frame = document.querySelector('.stage iframe');
  if (!frame || event.source !== frame.contentWindow) return;
  const key = event.data && event.data.chronoskinKey;
  if (typeof key === 'string' && /^(r|a|ArrowLeft|ArrowRight)$/.test(key)) press(key);
});
