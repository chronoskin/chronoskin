// The colour editor. Each row has a picker, a text field and a reset
// button; the text field is what the form sends, so a colour with alpha
// that the picker cannot show is still kept.
document.addEventListener('alpine:init', () => {
  Alpine.data('colours', () => ({
    open() {
      this.$refs.dialog.showModal();
    },

    close() {
      this.$refs.dialog.close();
    },

    pick(event) {
      const row = event.target.closest('.colour');
      row.querySelector('.value').value = event.target.value;
      this.mark(row);
    },

    type(event) {
      const row = event.target.closest('.colour');
      this.show(row, event.target.value);
    },

    reset(event) {
      const row = event.currentTarget.closest('.colour');
      row.querySelector('.value').value = row.dataset.own;
      this.show(row, row.dataset.own);
    },

    show(row, text) {
      const hex = toHex(text);
      if (hex) row.querySelector('.swatch').value = hex;
      this.mark(row);
    },

    mark(row) {
      row.classList.toggle('is-custom', row.querySelector('.value').value.trim() !== row.dataset.own);
    },
  }));
});

// Gives #rrggbb for #rgb, #rrggbb, rgb() or rgba(), dropping any alpha, or
// null when the text is not a colour yet.
function toHex(text) {
  text = text.trim();
  let m = /^#([0-9a-f]{3})[0-9a-f]?$/i.exec(text);
  if (m) return '#' + [...m[1]].map((c) => c + c).join('').toLowerCase();
  m = /^#([0-9a-f]{6})([0-9a-f]{2})?$/i.exec(text);
  if (m) return '#' + m[1].toLowerCase();
  m = /^rgba?\(\s*(\d{1,3})\s*[, ]\s*(\d{1,3})\s*[, ]\s*(\d{1,3})/i.exec(text);
  if (m && m.slice(1, 4).every((n) => Number(n) <= 255)) {
    return '#' + m.slice(1, 4).map((n) => Number(n).toString(16).padStart(2, '0')).join('');
  }
  return null;
}
