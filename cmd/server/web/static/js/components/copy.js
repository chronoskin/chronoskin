// A command with a Copy button. The whole command is copied even when the
// line shows it cut short.
document.addEventListener('alpine:init', () => {
  Alpine.data('copy', () => ({
    label: 'Copy',

    copy() {
      const text = this.$root.querySelector('code').textContent;
      navigator.clipboard.writeText(text).then(() => {
        this.label = 'Copied';
        setTimeout(() => { this.label = 'Copy'; }, 1500);
      });
    },
  }));
});
