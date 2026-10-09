// The heart on a style: saves or removes it. What is saved comes from the
// button's data attributes: id, title and the design width.
document.addEventListener('alpine:init', () => {
  Alpine.data('favourite', () => ({
    get saved() {
      return Alpine.store('saved').has(this.$root.dataset.id);
    },

    get pressed() {
      return this.saved ? 'true' : 'false';
    },

    get label() {
      return this.saved ? 'Remove from saved styles' : 'Save this style';
    },

    toggle() {
      const { id, title, width } = this.$root.dataset;
      Alpine.store('saved').toggle({ id, title, width: Number(width) });
    },
  }));
});
