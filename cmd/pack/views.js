// Walks the views of a specimen and reports what a visitor would trip over:
// links that lead nowhere, a view that cannot be left, a view that cannot
// be reached by clicking from the first one, text nearly the colour of its
// ground, and at phone width a page that scrolls sideways or text cut off
// at the edge. On a page prepared with every menu and tooltip held open
// (data-open on the root) it reports only those that leave the window or
// are clipped. Evaluates to a JSON list of sentences.
(async () => {
  const sleep = ms => new Promise(r => setTimeout(r, ms));
  // narrow: a tablet or a phone, where a responsive layout rearranges itself.
  const w = innerWidth, narrow = w <= 800, phone = narrow, out = [];
  const opened = document.documentElement.hasAttribute('data-open');
  // Measure the settled page, not a colour half-way through a transition.
  const still = document.createElement('style');
  still.textContent = '*, *::before, *::after { transition: none !important; animation: none !important; }';
  document.head.appendChild(still);
  const views = [...document.querySelectorAll('.ds-view')];
  const ids = views.map(v => v.id);
  const shown = () => views.filter(v => getComputedStyle(v).display !== 'none').map(v => v.id);
  const first = shown();
  if (first.length !== 1) out.push('without a fragment ' + first.length + ' views show');
  const hrefs = [...document.querySelectorAll('a[href^="#"]')].map(a => a.getAttribute('href').slice(1));
  if (!phone && !opened) {
    const dead = [...new Set(hrefs.filter(h => h && !document.getElementById(h)))];
    if (dead.length) out.push('links to nothing: #' + dead.join(' #'));
  }
  // A link counts when it can be seen and clicked inside the window, or
  // sits in a closed menu that opens on a click.
  const reachable = () => {
    const set = new Set();
    for (const a of document.querySelectorAll('a[href^="#"]')) {
      const r = a.getBoundingClientRect();
      let hidden = false, menu = false;
      for (let p = a; p; p = p.parentElement) {
        const cs = getComputedStyle(p);
        if (cs.display === 'none' || cs.visibility === 'hidden') { hidden = true; break; }
        if (p.tagName === 'DETAILS' && !p.open) { menu = true; break; }
      }
      if (hidden) continue;
      if (menu || (r.width > 1 && r.height > 1 && r.right > 4 && r.left < w - 4)) set.add(a.getAttribute('href').slice(1));
    }
    return set;
  };
  // Colours as the eye gets them: text over the fills under it.
  // A computed colour is "rgb(r, g, b)", "rgba(...)" or, when it was mixed,
  // "color(srgb r g b / a)" with parts from 0 to 1. Anything else is unknown.
  const rgba = c => {
    const m = c.match(/-?[\d.]+(e-?\d+)?/g);
    if (!m || m.length < 3) return null;
    if (c.startsWith('rgb')) return [+m[0], +m[1], +m[2], m.length > 3 ? +m[3] : 1];
    if (c.startsWith('color(srgb')) return [m[0] * 255, m[1] * 255, m[2] * 255, m.length > 3 ? +m[3] : 1];
    return null;
  };
  const lum = c => { const f = v => { v /= 255; return v <= 0.04045 ? v / 12.92 : Math.pow((v + 0.055) / 1.055, 2.4); }; return 0.2126 * f(c[0]) + 0.7152 * f(c[1]) + 0.0722 * f(c[2]); };
  const over = (top, under) => { const a = top[3]; return [top[0] * a + under[0] * (1 - a), top[1] * a + under[1] * (1 - a), top[2] * a + under[2] * (1 - a), 1]; };
  // The fill behind a piece of text, or null when a picture or a gradient
  // makes it unknowable. What lies behind is everything painted under the
  // middle of the text, not only its ancestors: a hero is often a sibling
  // pulled up under a clear bar, or a layer drawn by ::before.
  const painted = el => {
    for (const part of [null, '::before', '::after']) {
      const cs = getComputedStyle(el, part);
      if (part && (cs.content === 'none' || cs.display === 'none')) continue;
      if (cs.backgroundImage !== 'none') return null;
      if (part) { const c = rgba(cs.backgroundColor); if (c && c[3] > 0) return null; }
    }
    return rgba(getComputedStyle(el).backgroundColor);
  };
  const ground = el => {
    el.scrollIntoView({ block: 'center' });
    const r = el.getBoundingClientRect();
    const stack = document.elementsFromPoint(r.left + Math.min(r.width / 2, 40), r.top + r.height / 2);
    const from = stack.indexOf(el);
    if (from < 0) return null;
    // A table paints its rows and row groups under their cells, but they
    // are not found under a point: put them back where they are painted.
    const under0 = [];
    for (const item of stack.slice(from)) {
      under0.push(item);
      if (item.tagName === 'TD' || item.tagName === 'TH') {
        for (let p = item.parentElement; p && p.tagName !== 'TABLE'; p = p.parentElement) if (!under0.includes(p)) under0.push(p);
      }
    }
    const layers = [];
    for (const under of under0) {
      const c = painted(under);
      if (c === null) return null;
      if (c[3] > 0) { layers.push(c); if (c[3] >= 1) break; }
    }
    let bg = [255, 255, 255, 1];
    for (let i = layers.length - 1; i >= 0; i--) bg = over(layers[i], bg);
    return bg;
  };
  const ruler = document.createElement('canvas').getContext('2d');
  const graph = {};
  for (const v of ids) {
    location.hash = v;
    await sleep(50);
    const s = shown();
    if (s.length !== 1 || s[0] !== v) { out.push('#' + v + ' shows [' + s.join(', ') + ']'); continue; }
    graph[v] = ids.filter(x => x !== v && reachable().has(x));
    if (!graph[v].length && !opened) out.push('#' + v + ': no visible link to any other view');
    if (opened) {
      // Every menu and tooltip is held open: none may leave the window or
      // be cut by a box that clips it.
      for (const el of document.querySelectorAll('.ds-view:target [class*="ds-menu"], .ds-view:target [class*="ds-tooltip"], .ds-nav [class*="ds-menu"], header [class*="ds-menu"]')) {
        const cs = getComputedStyle(el);
        if (cs.position !== 'absolute' && cs.position !== 'fixed') continue;
        const r = el.getBoundingClientRect();
        if (r.width < 4 || r.height < 4 || cs.visibility === 'hidden' || +cs.opacity < 0.1) continue;
        const name = '.' + String(el.className).split(' ')[0];
        if (r.right > w + 2) out.push('#' + v + ': open ' + name + ' runs ' + Math.round(r.right - w) + 'px past the right edge');
        else if (r.left < -2) out.push('#' + v + ': open ' + name + ' runs ' + Math.round(-r.left) + 'px past the left edge');
        const holder = el.offsetParent;
        for (let p = el.parentElement; p && holder; p = p.parentElement) {
          const pc = getComputedStyle(p);
          if (pc.overflowX === 'visible' && pc.overflowY === 'visible') continue;
          if (p !== holder && !p.contains(holder)) continue;
          const pr = p.getBoundingClientRect();
          if (r.right > pr.right + 4 || r.left < pr.left - 4 || r.bottom > pr.bottom + 4 || r.top < pr.top - 4) { out.push('#' + v + ': open ' + name + ' is cut off by .' + String(p.className).split(' ')[0]); break; }
        }
      }
      continue;
    }
    let faint = 0, faintSample = '';
    for (const el of document.querySelectorAll('.ds-view:target *, .ds-nav *, footer *')) {
      if (![...el.childNodes].some(n => n.nodeType === 3 && n.textContent.trim().length > 1)) continue;
      const r = el.getBoundingClientRect(), cs = getComputedStyle(el);
      if (r.width < 2 || r.height < 2 || cs.visibility === 'hidden' || +cs.opacity < 0.5 || el.closest('[disabled], [aria-disabled="true"]')) continue;
      let hidden = false;
      for (let p = el.parentElement; p; p = p.parentElement) { const pc = getComputedStyle(p); if (pc.display === 'none' || pc.visibility === 'hidden' || +pc.opacity < 0.5 || (p.tagName === 'DETAILS' && !p.open && !el.closest('summary'))) { hidden = true; break; } }
      if (hidden) continue;
      const bg = ground(el), fg = rgba(cs.color);
      if (!bg || !fg || cs.webkitTextFillColor === 'rgba(0, 0, 0, 0)' || cs.backgroundClip === 'text') continue;
      const a = lum(over(fg, bg)), b = lum(bg);
      if ((Math.max(a, b) + 0.05) / (Math.min(a, b) + 0.05) < 1.6) { faint++; faintSample = faintSample || '.' + String(el.className || el.tagName).split(' ')[0] + ' "' + el.textContent.trim().slice(0, 24) + '"'; }
    }
    // A letter, a figure or a sign in a small square or round box (an
    // avatar, a step number, a count, a key) must sit in its middle: the
    // eye catches one pixel off. The ink is measured, not the line box.
    let off = 0, offSample = '';
    for (const el of document.querySelectorAll('.ds-view:target *, .ds-nav *')) {
      if (el.children.length || el.childNodes.length !== 1) continue;
      const cs = getComputedStyle(el);
      let text = el.textContent.trim();
      if (cs.textTransform === 'uppercase') text = text.toUpperCase();
      if (!/^[A-Z0-9+]{1,3}$/.test(text)) continue;
      const r = el.getBoundingClientRect();
      if (r.width < 14 || r.height < 14 || r.width > 72 || r.height > 72 || r.width / r.height < 0.75 || r.width / r.height > 1.34) continue;
      const boxed = cs.backgroundImage !== 'none' || (rgba(cs.backgroundColor) || [0, 0, 0, 0])[3] > 0 || parseFloat(cs.borderTopWidth) > 0;
      if (!boxed || cs.visibility === 'hidden' || cs.display.startsWith('table')) continue;
      // Only a box that means to centre its sign: a right-aligned figure in a cell is not one.
      const centred = cs.textAlign === 'center' || (/flex|grid/.test(cs.display) && /center/.test(cs.justifyContent + cs.justifyItems + cs.placeItems));
      if (!centred) continue;
      const range = document.createRange();
      range.selectNodeContents(el);
      const line = range.getClientRects()[0];
      if (!line) continue;
      ruler.font = cs.fontStyle + ' ' + cs.fontWeight + ' ' + cs.fontSize + ' ' + cs.fontFamily;
      const m = ruler.measureText(text);
      const base = line.top + (line.height - (m.fontBoundingBoxAscent + m.fontBoundingBoxDescent)) / 2 + m.fontBoundingBoxAscent;
      const dy = (base - m.actualBoundingBoxAscent + base + m.actualBoundingBoxDescent) / 2 - (r.top + r.bottom) / 2;
      const spaced = parseFloat(cs.letterSpacing) || 0;
      const dx = spaced ? 0 : line.left + (m.actualBoundingBoxRight - m.actualBoundingBoxLeft) / 2 - (r.left + r.right) / 2;
      if (Math.abs(dy) > Math.max(1.5, r.height * 0.07) || Math.abs(dx) > Math.max(1.5, r.width * 0.07)) {
        off++;
        offSample = offSample || '.' + String(el.className || el.tagName).split(' ')[0] + ' "' + text + '" (' + (Math.abs(dy) >= Math.abs(dx) ? Math.round(Math.abs(dy) * 10) / 10 + 'px too ' + (dy > 0 ? 'low' : 'high') : Math.round(Math.abs(dx) * 10) / 10 + 'px too far ' + (dx > 0 ? 'right' : 'left')) + ')';
      }
    }
    if (off) out.push('#' + v + ': ' + off + ' small boxes with their sign off centre, such as ' + offSample);
    if (faint) out.push('#' + v + ': ' + faint + ' pieces of text nearly the colour of their ground, such as ' + faintSample);
    if (phone) {
      // What the bar offers on a wide screen must still be offered here:
      // a link the phone rules hide needs another link, or a menu entry,
      // to the same place.
      const offered = reachable(), lostLinks = [];
      for (const a of document.querySelectorAll('.ds-nav a[href^="#"], header a[href^="#"]')) {
        const to = a.getAttribute('href').slice(1);
        if (!to || offered.has(to) || !document.getElementById(to)) continue;
        const label = (a.textContent.trim() || a.getAttribute('aria-label') || to).slice(0, 24);
        if (!lostLinks.includes(label)) lostLinks.push(label);
      }
      if (lostLinks.length) out.push('#' + v + ': the navigation hides ' + lostLinks.map(l => '"' + l + '"').join(', ') + ' and offers no other link to the same place');
    }
    if (!phone) continue;
    const sw = document.documentElement.scrollWidth;
    if (sw > w + 1) out.push('#' + v + ': the page scrolls sideways (' + sw + 'px wide)');
    let cut = 0, sample = '';
    for (const el of document.querySelectorAll('.ds-view:target *')) {
      if (el.children.length || !el.textContent.trim()) continue;
      const r = el.getBoundingClientRect();
      if (r.width < 2 || r.height < 2 || r.right <= w + 2 || r.left >= w) continue;
      let scroller = false;
      for (let p = el.parentElement; p; p = p.parentElement) {
        const ox = getComputedStyle(p).overflowX;
        if (ox === 'auto' || ox === 'scroll') { scroller = true; break; }
      }
      if (!scroller) { cut++; sample = sample || '.' + String(el.className || el.tagName).split(' ')[0]; }
    }
    if (cut) out.push('#' + v + ': ' + cut + ' pieces of text cut off at the right edge, such as ' + sample);
  }
  if (opened && !phone) {
    // Rules that style nothing in the specimen and name a class STYLE.md
    // does not describe are leftovers of earlier work. States are taken out
    // of a selector before it is tried, and the ".is-" classes that show a
    // state in the specimen are let be.
    const states = /::?(before|after|placeholder|marker|selection|first-letter|first-line|backdrop|-webkit-[a-z-]+|-moz-[a-z-]+)|:(hover|focus-visible|focus-within|focus|active|checked|disabled|enabled|target|visited|link|any-link|placeholder-shown|invalid|valid|required|optional|indeterminate|default|read-only|open)\b|\[open\]/g;
    const documented = new Set((document.documentElement.dataset.documented || '').split(' '));
    const dead = [];
    const walk = rules => {
      for (const rule of rules) {
        if (rule.cssRules && !rule.selectorText) { walk(rule.cssRules); continue; }
        if (!rule.selectorText) continue;
        const live = rule.selectorText.split(/,(?![^()]*\))/).some(sel => {
          if (/\.is-|:root|^\s*(html|body)\b/.test(sel)) return true;
          try { return document.querySelector(sel.replace(states, '').replace(/:not\(\s*\)|:is\(\s*\)|:where\(\s*\)|:has\(\s*\)/g, '') || '*') !== null; } catch { return true; }
        });
        if (live) continue;
        // A class that the specimen does not show but STYLE.md describes is
        // part of the kit, to build with; one that neither knows is left over.
        const unknown = (rule.selectorText.match(/\.ds-[a-z0-9_-]+/g) || []).filter(c => !documented.has(c) && !document.querySelector(c));
        if (unknown.length) { dead.push(unknown[0]); if (location.search === '?list') out.push('leftover rule: ' + rule.selectorText); }
        if (rule.cssRules) walk(rule.cssRules);
      }
    };
    for (const sheet of document.styleSheets) { try { walk(sheet.cssRules); } catch { /* a sheet that cannot be read is not judged */ } }
    const classes = [...new Set(dead)];
    if (classes.length) out.push(dead.length + ' rules for ' + classes.length + ' classes that neither the specimen nor STYLE.md has, such as ' + classes.slice(0, 5).join(' '));
  }
  const seen = new Set([ids[0]]), queue = [ids[0]];
  while (queue.length) for (const x of graph[queue.shift()] || []) if (!seen.has(x)) { seen.add(x); queue.push(x); }
  const lost = ids.filter(x => !seen.has(x) && graph[x]);
  if (lost.length && !opened) out.push('cannot be reached by clicking from #' + ids[0] + ': ' + lost.join(', '));
  return JSON.stringify(out);
})()
