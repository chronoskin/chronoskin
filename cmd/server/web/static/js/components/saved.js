// The saved page: the list of saved styles and its filter by era.
document.addEventListener('alpine:init', () => {
  Alpine.data('savedList', () => ({
    // The era chosen in the filter, as the code a short ID carries
    // ("v1-yk-3333" is of era "yk"); empty for every era.
    era: '',
    // False from the start: the placeholder cards show only until this runs.
    waiting: false,

    init() {
      // Only eras that have a saved style are offered.
      Alpine.effect(() => {
        const codes = new Set(this.items.map((item) => item.era));
        for (const option of this.$root.querySelectorAll('.cselect .menu label')) {
          const value = option.querySelector('input').value;
          option.classList.toggle('absent', value !== '' && !codes.has(value));
        }
        if (this.era && !codes.has(this.era)) this.all();
      });
    },

    get items() {
      return Alpine.store('saved').items.map((item) => ({
        id: item.id,
        era: item.id.split('-')[1] || '',
        // Styles saved by an earlier version carry "era" instead of "title".
        title: item.title || item.era || 'Saved style',
        width: item.width || 1280,
        url: '/s/' + item.id,
        specimen: '/s/' + item.id + '/specimen.html?frame',
      }));
    },

    get shown() {
      return this.era ? this.items.filter((item) => item.era === this.era) : this.items;
    },

    get empty() {
      return Alpine.store('saved').items.length === 0;
    },

    get some() {
      return !this.empty;
    },

    get count() {
      const all = this.items.length;
      return (this.era ? this.shown.length + ' of ' + all : all) + ' saved';
    },

    pick(event) {
      if (event.target.name === 'saved-era') this.era = event.target.value;
    },

    all() {
      this.era = '';
      this.$root.querySelector('.cselect input[type="radio"]').checked = true;
    },
  }));
});
