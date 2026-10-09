// One saved style's card. Removing it is asked in the card first.
document.addEventListener('alpine:init', () => {
  Alpine.data('savedCard', () => ({
    asking: false,

    get state() { return this.asking ? 'is-asking' : ''; },

    ask() { this.asking = true; },
    keep() { this.asking = false; },

    remove() {
      Alpine.store('saved').remove(this.$root.dataset.id);
    },
  }));
});
