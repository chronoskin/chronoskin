// The footer: marks itself "is-long" on a page long enough to scroll, where
// it offers a way back to the top.
document.addEventListener('alpine:init', () => {
  Alpine.data('foot', () => ({
    long: false,

    init() {
      const measure = () => {
        this.long = document.documentElement.scrollHeight > window.innerHeight + 1;
      };
      // Pictures load and lists are filtered: the page's height changes.
      new ResizeObserver(measure).observe(document.body);
      window.addEventListener('resize', measure);
      measure();
    },

    get state() { return this.long ? 'is-long' : ''; },

    top() {
      const calm = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
      window.scrollTo({ top: 0, behavior: calm ? 'auto' : 'smooth' });
    },
  }));
});
