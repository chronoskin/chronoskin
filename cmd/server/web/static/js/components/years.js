// The year filter of the home page: two sliders on one track. An era shows
// when its years overlap the chosen span; one that does not is also taken
// out of the choice.
document.addEventListener('alpine:init', () => {
  Alpine.data('years', () => ({
    from: 0,
    to: 0,
    shown: 0,
    cards: [],

    init() {
      this.cards = [...this.$root.querySelectorAll('.cards > .card[data-from]')];
      this.show(this.first(), this.last());
    },

    get count() { return this.shown === 1 ? '1 era' : this.shown + ' eras'; },
    get narrowed() { return this.from !== this.first() || this.to !== this.last(); },

    first() { return Number(this.$root.dataset.first); },
    last() { return Number(this.$root.dataset.last); },

    // The two handles may not cross: the one being moved stops at the other.
    move(event) {
      let from = Number(this.$refs.from.value);
      let to = Number(this.$refs.to.value);
      if (from > to) {
        if (event.target === this.$refs.from) from = to;
        else to = from;
      }
      this.show(from, to);
    },

    clear() {
      for (const card of this.cards) card.querySelector('input').checked = false;
    },

    // With the years narrowed and no era chosen, Generate picks among the
    // eras that show, not among all of them.
    fill(event) {
      if (!this.narrowed || event.formData.getAll('era').length) return;
      for (const card of this.cards) {
        if (!card.hidden) event.formData.append('era', card.querySelector('input').value);
      }
    },

    show(from, to) {
      this.from = from;
      this.to = to;
      this.$refs.from.value = from;
      this.$refs.to.value = to;
      const span = this.last() - this.first() || 1;
      // 9px is half a handle: the fill ends at the handles' centres.
      const at = (year) => 'calc(9px + (100% - 18px) * ' + (year - this.first()) / span + ')';
      this.$refs.fill.style.left = at(from);
      this.$refs.fill.style.right = 'calc(100% - ' + at(to) + ')';
      let shown = 0;
      for (const card of this.cards) {
        card.hidden = Number(card.dataset.from) > to || Number(card.dataset.to) < from;
        const box = card.querySelector('input');
        box.disabled = card.hidden;
        if (card.hidden) box.checked = false;
        else shown++;
      }
      this.shown = shown;
    },
  }));
});
