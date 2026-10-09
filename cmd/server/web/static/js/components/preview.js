// Shows a specimen at a chosen width, scaled down when the stage is
// narrower, so a 1440px design is seen whole. The width starts as the
// style's own (data-width); three buttons switch it to tablet and phone,
// and the choice is kept for the visit. A layout without data-responsive
// has one width: a tablet or phone shows it whole and scaled down.
// The stage is x-ref="stage", or the element itself when there are no
// buttons, as on the saved page.
const TABLET = 768;
const PHONE = 390;
const DEVICE = 'chronoskin-device';

function remembered() {
  try {
    const width = Number(sessionStorage.getItem(DEVICE));
    return width === TABLET || width === PHONE ? width : 0;
  } catch {
    return 0;
  }
}

function remember(width) {
  try {
    sessionStorage.setItem(DEVICE, String(width));
  } catch {
    // The choice lasts until the next page.
  }
}

document.addEventListener('alpine:init', () => {
  Alpine.data('preview', () => ({
    // The chosen device width; 0 means the style's own.
    chosen: 0,
    stage: null,

    init() {
      this.stage = this.$refs.stage || this.$root;
      if (this.$refs.stage) this.chosen = remembered();
      new ResizeObserver(() => this.fit()).observe(this.stage);
      // The saved page binds data-width after this runs.
      new MutationObserver(() => this.fit()).observe(this.$root, { attributes: true, attributeFilter: ['data-width'] });
      this.fit();
    },

    get isDesktop() { return this.chosen === 0 ? 'true' : 'false'; },
    get isTablet() { return this.chosen === TABLET ? 'true' : 'false'; },
    get isPhone() { return this.chosen === PHONE ? 'true' : 'false'; },

    desktop() { this.show(0); },
    tablet() { this.show(TABLET); },
    phone() { this.show(PHONE); },

    show(width) {
      this.chosen = width;
      remember(width);
      this.fit();
    },

    fit() {
      const frame = this.stage.querySelector('iframe');
      const own = Number(this.$root.dataset.width) || 1280;
      // "device" is the window the specimen is seen through, "width" what
      // it is laid out at: the same, unless the layout has only one width.
      const device = this.chosen || own;
      const width = this.chosen && !('responsive' in this.$root.dataset) ? Math.max(own, device) : device;
      const scale = Math.min(1, this.stage.clientWidth / device) * (device / width);
      frame.style.width = width + 'px';
      frame.style.height = this.stage.clientHeight / scale + 'px';
      frame.style.transform = 'scale(' + scale + ')';
    },
  }));
});
