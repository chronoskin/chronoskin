// The list of sections beside a document: marks the section being read
// and slides a marker to it.
document.addEventListener('alpine:init', () => {
  Alpine.data('toc', () => ({
    links: [],
    sections: [],

    init() {
      this.links = [...this.$root.querySelectorAll('.toc a')];
      this.sections = this.links.map((link) => document.getElementById(link.hash.slice(1)));
      this.update = this.update.bind(this);
      window.addEventListener('scroll', this.update, { passive: true });
      window.addEventListener('resize', this.update);
      this.update();
    },

    destroy() {
      window.removeEventListener('scroll', this.update);
      window.removeEventListener('resize', this.update);
    },

    // The current section is the last one whose top has passed a line a
    // third of the way down the window; at the end of the page, the last.
    update() {
      const line = window.innerHeight / 3;
      let current = 0;
      this.sections.forEach((section, i) => {
        if (section.getBoundingClientRect().top <= line) current = i;
      });
      if (window.innerHeight + window.scrollY >= document.documentElement.scrollHeight - 2) current = this.sections.length - 1;
      this.links.forEach((link, i) => {
        if (i === current) link.setAttribute('aria-current', 'true');
        else link.removeAttribute('aria-current');
      });
      const link = this.links[current];
      const marker = this.$refs.marker;
      marker.style.transform = 'translateY(' + link.offsetTop + 'px)';
      marker.style.height = link.offsetHeight + 'px';
    },
  }));
});
