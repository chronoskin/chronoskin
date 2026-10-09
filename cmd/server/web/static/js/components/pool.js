// The eras and page types chosen on a style page. When a choice changes
// and its list closes, the address and the pieces of the page that carry
// the choice are replaced with fresh ones from the server. The specimen is
// left alone, so it does not flash.
const FOLLOWS_CHOICE = ['.parts', '.pane.link code', '.actions [data-key="a"]', '.editor form > input[type="hidden"]', '.editor footer'];

document.addEventListener('alpine:init', () => {
  Alpine.data('pool', () => ({
    async apply() {
      const checked = (name) => [...this.$root.querySelectorAll('input[name="' + name + '"]:checked')].map((box) => box.value);
      const query = new URLSearchParams();
      const keep = checked('lock');
      if (keep.length) query.set('keep', keep.join(','));
      const types = checked('archetype');
      if (types.length) query.set('type', types.join(','));
      const eras = checked('era');
      if (eras.length) query.set('eras', eras.join(','));
      const next = query.toString();
      const url = location.pathname + (next ? '?' + next : '');
      try {
        const response = await fetch(url, { headers: { Accept: 'text/html' } });
        if (!response.ok) throw new Error(response.status);
        const fresh = new DOMParser().parseFromString(await response.text(), 'text/html');
        for (const selector of FOLLOWS_CHOICE) {
          const now = document.querySelectorAll(selector);
          const then = fresh.querySelectorAll(selector);
          if (now.length !== then.length) throw new Error('the page changed shape');
          now.forEach((element, i) => element.replaceWith(document.importNode(then[i], true)));
        }
        history.replaceState(null, '', url);
        // GoatCounter binds to the elements it found at load.
        if (window.goatcounter && window.goatcounter.bind_events) window.goatcounter.bind_events();
      } catch {
        location.assign(url);
      }
    },
  }));
});
