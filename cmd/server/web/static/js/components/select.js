// The custom select is a <details> of radio buttons or checkboxes and works
// on its own. This closes it after a choice, on Escape and on a click
// outside, filters the options by what is typed, submits the form when the
// select is marked data-submit, and dispatches "select-done" when a list
// closes with its choice changed.
document.addEventListener('alpine:init', () => {
  Alpine.data('select', () => ({
    chosen: 0,
    dirty: false,

    init() {
      this.tally();
      // A form reset clears the boxes without a change event.
      const form = this.$root.closest('form');
      if (form) form.addEventListener('reset', () => setTimeout(() => this.tally()));
      this.$root.addEventListener('toggle', () => {
        if (this.$root.open) {
          // The list reaches the foot of the window at most.
          const menu = this.$root.querySelector('.menu');
          const room = window.innerHeight - menu.getBoundingClientRect().top - 12;
          menu.style.maxHeight = Math.max(180, room) + 'px';
          this.$root.querySelector('.filter').focus();
        } else {
          this.clear();
          this.done();
        }
      });
    },

    get many() { return this.chosen > 1; },
    get some() { return this.chosen > 0; },
    get nothing() { return this.chosen === 0; },
    get count() { return String(this.chosen); },

    tally() {
      this.chosen = this.$root.querySelectorAll('.menu input:checked').length;
    },

    none() {
      for (const box of this.$root.querySelectorAll('.menu input:checked')) box.checked = false;
      this.tally();
      this.dirty = true;
      if (!this.$root.open) this.done();
    },

    done() {
      if (!this.dirty) return;
      this.dirty = false;
      this.$dispatch('select-done');
    },

    close() {
      this.$root.open = false;
    },

    choose(event) {
      this.tally();
      this.dirty = true;
      if (event.target.type !== 'radio') return;
      this.close();
      this.$root.querySelector('summary').focus();
      if ('submit' in this.$root.dataset) this.$root.closest('form').submit();
    },

    escape() {
      if (!this.$root.open) return;
      this.close();
      this.$root.querySelector('summary').focus();
    },

    filter(event) {
      const wanted = event.target.value.trim().toLowerCase();
      for (const option of this.$root.querySelectorAll('.menu label')) {
        option.hidden = !option.textContent.toLowerCase().includes(wanted);
      }
    },

    clear() {
      this.$root.querySelector('.filter').value = '';
      for (const option of this.$root.querySelectorAll('.menu label')) option.hidden = false;
    },
  }));
});
