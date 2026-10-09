// Saved styles live in this browser only, under one localStorage key.
// Storage can be missing or blocked, so every access is guarded.
const SAVED_KEY = 'chronoskin.saved';

function readSaved() {
  try {
    const items = JSON.parse(localStorage.getItem(SAVED_KEY) || '[]');
    return Array.isArray(items) ? items.filter((item) => item && typeof item.id === 'string') : [];
  } catch {
    return [];
  }
}

document.addEventListener('alpine:init', () => {
  Alpine.store('saved', {
    items: readSaved(),

    init() {
      // Another tab may save or remove a style.
      window.addEventListener('storage', (event) => {
        if (event.key === SAVED_KEY) this.items = readSaved();
      });
    },

    has(id) {
      return this.items.some((item) => item.id === id);
    },

    toggle(item) {
      this.items = this.has(item.id) ? this.items.filter((other) => other.id !== item.id) : [item, ...this.items];
      this.write();
    },

    remove(id) {
      this.items = this.items.filter((item) => item.id !== id);
      this.write();
    },

    write() {
      try {
        localStorage.setItem(SAVED_KEY, JSON.stringify(this.items));
      } catch {
        // The list still works until the page is closed.
      }
    },
  });
});
