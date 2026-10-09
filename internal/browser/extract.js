// Measures the design of the current page from computed styles and returns a
// JSON string. Lists are weighted by painted area (backgrounds, layout) or by
// characters (text). `metrics` holds the single numbers that Never rules and
// the feature vector are computed from.
(() => {
  // A specimen shows one view at a time; measure all of them.
  // They go back to one at a time before this returns, so that a screenshot
  // taken afterwards shows the page as a visitor sees it.
  const views = [...document.querySelectorAll('.ds-view')];
  for (const view of views) view.style.setProperty('display', 'block', 'important');
  const add = (m, k, v) => { m[k] = (m[k] || 0) + v; };
  const top = (m, n) => Object.entries(m).sort((a, b) => b[1] - a[1]).slice(0, n)
    .map(([k, v]) => ({ value: k, weight: Math.round(v) }));
  const px = (v) => parseFloat(v) || 0;
  const median = (a) => { if (!a.length) return 0; const s = [...a].sort((x, y) => x - y); return s[Math.floor(s.length / 2)]; };
  const round1 = (n) => Math.round(n * 10) / 10;
  const quantile = (a, q) => { if (!a.length) return 0; const s = [...a].sort((x, y) => x - y); return s[Math.min(s.length - 1, Math.floor(s.length * q))]; };

  const bg = {}, text = {}, fonts = {}, families = {}, sizes = {}, links = {}, borders = {}, borderColours = {},
    radii = {}, shadows = {}, textShadows = {}, gradients = {}, layout = {}, transforms = {}, lineRatios = {}, sizeChars = {};
  const deprecated = {};
  let boxes = 0, bordered = 0, rounded = 0, roundShapes = 0, shadowed = 0, gradientBoxes = 0, bgImages = 0;
  let chars = 0, linkChars = 0, underlined = 0, upperChars = 0, boldChars = 0, textShadowChars = 0;
  let maxRadius = 0, maxBorder = 0, maxBlur = 0, maxOffset = 0, maxTracking = 0;
  let minSize = Infinity, maxSize = 0, minWeight = Infinity, maxWeight = 0;
  let transitions = 0, animations = 0, mediaArea = 0, paintedArea = 0, backdropBlur = 0, translucent = 0, paintedBoxes = 0;
  const borderWidths = [], radiusValues = [], textEdges = [];
  const viewportWidth = document.documentElement.clientWidth;

  for (const el of document.querySelectorAll('body *')) {
    const tag = el.tagName.toLowerCase();
    if (['font', 'center', 'marquee', 'frame', 'frameset', 'blink'].includes(tag)) add(deprecated, tag, 1);
    const r = el.getBoundingClientRect();
    if (r.width < 2 || r.height < 2) continue;
    const s = getComputedStyle(el);
    if (s.visibility === 'hidden' || s.display === 'none') continue;
    boxes++;
    const area = r.width * r.height;
    if (['img', 'svg', 'canvas', 'video', 'picture'].includes(tag)) mediaArea += area;
    if (s.backgroundColor !== 'rgba(0, 0, 0, 0)') {
      paintedBoxes++;
      const alpha = s.backgroundColor.match(/rgba\([^)]*,\s*([\d.]+)\)|\/\s*([\d.]+)\)/);
      const opacity = alpha ? parseFloat(alpha[1] || alpha[2]) : 1;
      if (opacity < 0.95) translucent++;
      // A see-through fill is not the colour of its box: only fills that
      // are mostly opaque count toward the colours painted.
      if (opacity >= 0.5) { add(bg, s.backgroundColor, area); paintedArea += area; }
    }
    if ((s.backdropFilter || 'none') !== 'none') backdropBlur++;
    if (s.backgroundImage.includes('gradient')) { gradientBoxes++; add(gradients, s.backgroundImage.slice(0, 160), 1); }
    else if (s.backgroundImage !== 'none') bgImages++;
    const d = s.display;
    const kind = d.includes('table') ? 'table' : d.includes('flex') ? 'flex' : d.includes('grid') ? 'grid'
      : s.float !== 'none' ? 'float' : s.position === 'absolute' || s.position === 'fixed' ? 'absolute' : null;
    if (kind) add(layout, kind, area);
    if (s.transitionDuration.split(',').some((t) => parseFloat(t) > 0)) transitions++;
    if (s.animationName !== 'none') animations++;

    let own = 0;
    for (const n of el.childNodes) if (n.nodeType === 3) own += n.textContent.trim().length;
    if (own) {
      chars += own;
      const size = px(s.fontSize), weight = parseInt(s.fontWeight, 10) || 400;
      add(text, s.color, own);
      add(fonts, s.fontFamily, own);
      add(families, s.fontFamily.split(',')[0].trim().replace(/["']/g, '').toLowerCase(), own);
      add(sizes, `${s.fontSize} / ${s.lineHeight} / ${s.fontWeight}`, own);
      add(sizeChars, size, own);
      if (own >= 2) {
        minSize = Math.min(minSize, size); maxSize = Math.max(maxSize, size);
        minWeight = Math.min(minWeight, weight); maxWeight = Math.max(maxWeight, weight);
      }
      if (weight >= 600) boldChars += own;
      if (s.lineHeight !== 'normal' && size) add(lineRatios, `${size}|${round1(px(s.lineHeight) / size * 10) / 10}`, own);
      if (s.letterSpacing !== 'normal') maxTracking = Math.max(maxTracking, Math.abs(px(s.letterSpacing)));
      if (s.textTransform === 'uppercase') upperChars += own;
      if (s.textTransform !== 'none') add(transforms, s.textTransform, own);
      if (s.textShadow !== 'none') { add(textShadows, s.textShadow, own); textShadowChars += own; }
      textEdges.push([r.left, r.right, own]);
      const a = el.closest('a');
      if (a) {
        linkChars += own;
        const line = s.textDecorationLine !== 'none' ? s.textDecorationLine : getComputedStyle(a).textDecorationLine;
        if (line.includes('underline')) underlined += own;
        add(links, `${s.color} ${line}`, own);
      }
    }

    const widths = [s.borderTopWidth, s.borderRightWidth, s.borderBottomWidth, s.borderLeftWidth];
    const styles = [s.borderTopStyle, s.borderRightStyle, s.borderBottomStyle, s.borderLeftStyle];
    let hasBorder = false;
    widths.forEach((w, i) => {
      if (px(w) > 0 && styles[i] !== 'none' && styles[i] !== 'hidden') {
        hasBorder = true;
        maxBorder = Math.max(maxBorder, px(w));
        borderWidths.push(px(w));
      }
    });
    if (hasBorder) {
      bordered++;
      const i = widths.findIndex((w, j) => px(w) > 0 && styles[j] !== 'none');
      add(borders, `${widths[i]} ${styles[i]}`, 1);
      add(borderColours, [s.borderTopColor, s.borderRightColor, s.borderBottomColor, s.borderLeftColor][i], 1);
    }

    // A box rounded to half its shorter side is a pill or a circle, which
    // is a shape rather than a corner radius.
    const corner = s.borderTopLeftRadius;
    if (corner !== '0px') {
      const value = corner.endsWith('%') ? px(corner) / 100 * Math.min(r.width, r.height) : px(corner);
      if (value * 2 >= Math.min(r.width, r.height) - 1) roundShapes++;
      else { rounded++; add(radii, corner, 1); maxRadius = Math.max(maxRadius, value); radiusValues.push(value); }
    }
    if (s.boxShadow !== 'none') {
      shadowed++;
      add(shadows, s.boxShadow, 1);
      for (const one of s.boxShadow.split(/,(?![^(]*\))/)) {
        const n = (one.replace(/rgba?\([^)]*\)|color\([^)]*\)/g, '').match(/-?[\d.]+px/g) || []).map(px);
        maxOffset = Math.max(maxOffset, Math.abs(n[0] || 0), Math.abs(n[1] || 0));
        maxBlur = Math.max(maxBlur, n[2] || 0);
      }
    }
  }

  // Gaps between consecutive rows of tables and lists, and between
  // consecutive wide blocks that share a parent.
  const gaps = (els) => {
    const out = [];
    for (const el of els) {
      const next = el.nextElementSibling;
      if (!next || next.tagName !== el.tagName) continue;
      const a = el.getBoundingClientRect(), b = next.getBoundingClientRect();
      if (a.height < 2 || b.height < 2 || b.top < a.bottom - 1) continue;
      out.push(Math.max(0, b.top - a.bottom));
    }
    return out;
  };
  const rowGaps = gaps(document.querySelectorAll('tr, li'));
  const blockGaps = [];
  for (const parent of document.querySelectorAll('body, body *')) {
    const kids = [...parent.children].map((k) => k.getBoundingClientRect())
      .filter((k) => k.width >= viewportWidth * 0.4 && k.height >= 16);
    for (let i = 0; i + 1 < kids.length; i++) {
      const g = kids[i + 1].top - kids[i].bottom;
      if (g >= 0) blockGaps.push(g);
    }
  }

  // Content width: the span that holds the middle 90% of the text.
  let contentWidth = 0;
  if (chars) {
    const at = (idx, q) => {
      const s = [...textEdges].sort((a, b) => a[idx] - b[idx]);
      let seen = 0;
      for (const e of s) { seen += e[2]; if (seen >= chars * q) return e[idx]; }
      return s[s.length - 1][idx];
    };
    contentWidth = Math.max(0, at(1, 0.95) - at(0, 0.05));
  }

  const colours = new Set();
  for (const [c, a] of Object.entries(bg)) if (a >= paintedArea * 0.005) colours.add(c);
  for (const [c, n] of Object.entries(text)) if (n >= chars * 0.01) colours.add(c);
  for (const c of Object.keys(borderColours)) colours.add(c);

  const baseSize = Number((Object.entries(sizeChars).sort((a, b) => b[1] - a[1])[0] || [0])[0]);
  let lineHeight = 0, lineWeight = 0;
  for (const [k, n] of Object.entries(lineRatios)) {
    const [size, ratio] = k.split('|').map(Number);
    if (size === baseSize && n > lineWeight) { lineHeight = ratio; lineWeight = n; }
  }
  const familyCount = Object.values(families).filter((n) => n >= chars * 0.01).length;
  const body = getComputedStyle(document.body);
  const docW = document.documentElement.scrollWidth, docH = document.documentElement.scrollHeight;
  const pct = (n, of) => of ? round1(100 * n / of) : 0;
  const share = (m, k) => { const total = Object.values(m).reduce((a, b) => a + b, 0); return total ? round1(100 * (m[k] || 0) / total) : 0; };

  // Signals for the page type. They are counts, not a verdict: the
  // archetype is decided from them in Go.
  const bodyText = document.body.innerText || '';
  const count = (re) => (bodyText.match(re) || []).length;
  let numericRows = 0;
  for (const tr of document.querySelectorAll('tr')) {
    const cells = [...tr.children];
    if (cells.length >= 3 && tr.querySelector('a') && cells.filter((c) => /^[\d,.\s]+$/.test(c.textContent.trim()) && c.textContent.trim()).length >= 2) numericRows++;
  }
  let listDepth = 0;
  for (const ul of document.querySelectorAll('nav ul, aside ul, [class*="sidebar"] ul, [class*="toc"] ul')) {
    let d = 1, p = ul.parentElement;
    while (p) { if (p.tagName === 'UL' || p.tagName === 'OL') d++; p = p.parentElement; }
    listDepth = Math.max(listDepth, d);
  }
  let paragraphChars = 0;
  for (const p of document.querySelectorAll('p')) paragraphChars += p.textContent.trim().length;
  const hints = {
    numericRows,
    forumWords: count(/\b(forum|topics?|posts?|replies|threads?|moderators?)\b/gi),
    priceCount: count(/[$€£]\s?\d[\d,.]*/g),
    cartWords: count(/\b(add to (cart|basket)|checkout|in stock|free shipping)\b/gi),
    blogWords: count(/\b(posted (by|on|at)|permalink|comments?|archives?|trackback)\b/gi),
    dateCount: count(/\b(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)[a-z]*\.? \d{1,2}\b|\b\d{1,2} (jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)[a-z]*\b/gi),
    linkCount: document.links.length,
    linkTextPct: pct(linkChars, chars),
    paragraphTextPct: pct(Math.min(paragraphChars, chars), chars),
    codeBlocks: document.querySelectorAll('pre, code').length,
    navListDepth: listDepth,
    formControls: document.querySelectorAll('input:not([type=hidden]), select, textarea').length,
    passwordInputs: document.querySelectorAll('input[type=password]').length,
    buttons: document.querySelectorAll('button, input[type=submit], input[type=button]').length,
    signupWords: count(/\b(sign up|get started|start (your )?free|try (it )?(for )?free|download|pricing|free trial)\b/gi),
  };

  // Old pages that load Prototype.js define Array.prototype.toJSON, which
  // makes JSON.stringify encode every array as a string. Put it aside.
  const arrayToJSON = Array.prototype.toJSON;
  delete Array.prototype.toJSON;
  try {
  const result = JSON.stringify({
    title: document.title.slice(0, 80),
    url: location.href,
    textSample: (document.body.innerText || '').replace(/\s+/g, ' ').trim().slice(0, 1500),
    documentWidth: docW,
    documentHeight: docH,
    bodyBackground: body.backgroundColor,
    boxes, textChars: chars,
    backgroundsByArea: top(bg, 10),
    textColours: top(text, 8),
    fonts: top(fonts, 5),
    sizesByChars: top(sizes, 12),
    textTransforms: top(transforms, 3),
    links: top(links, 6),
    underlinedLinksPct: pct(underlined, linkChars),
    borders: top(borders, 6), borderColours: top(borderColours, 6), borderedBoxesPct: pct(bordered, boxes),
    radii: top(radii, 8), roundedBoxesPct: pct(rounded, boxes),
    boxShadows: top(shadows, 6), shadowedBoxesPct: pct(shadowed, boxes),
    textShadows: top(textShadows, 3),
    cssGradients: top(gradients, 5), gradientBoxesPct: pct(gradientBoxes, boxes),
    backgroundImageBoxes: bgImages,
    layoutByArea: top(layout, 5),
    deprecatedTags: deprecated,
    tables: document.querySelectorAll('table').length,
    images: document.images.length,
    archetypeHints: hints,
    metrics: {
      backdropBlurPct: pct(backdropBlur, boxes), translucentFillsPct: pct(translucent, paintedBoxes),
      viewportWidth,
      borderRadiusMax: round1(maxRadius), borderRadiusMin: round1(radiusValues.reduce((a, b) => Math.min(a, b), maxRadius)), borderRadiusMedian: round1(median(radiusValues)),
      roundedBoxesPct: pct(rounded, boxes), roundShapesPct: pct(roundShapes, boxes),
      borderWidthMax: round1(maxBorder), borderWidthMin: round1(borderWidths.reduce((a, b) => Math.min(a, b), maxBorder)), borderWidthMedian: round1(median(borderWidths)), borderedBoxesPct: pct(bordered, boxes),
      boxShadowPct: pct(shadowed, boxes), boxShadowBlurMax: round1(maxBlur), boxShadowOffsetMax: round1(maxOffset),
      textShadowPct: pct(textShadowChars, chars),
      gradientFillsPct: pct(gradientBoxes, boxes),
      fontSizeMin: chars ? round1(minSize) : 0, fontSizeMax: round1(maxSize), fontSizeBase: baseSize,
      fontWeightMin: chars ? minWeight : 0, fontWeightMax: maxWeight, boldTextPct: pct(boldChars, chars),
      fontFamilies: familyCount, fontFamilyMain: (top(families, 1)[0] || {}).value || '',
      lineHeight,
      underlinedLinksPct: pct(underlined, linkChars),
      letterSpacingMax: round1(maxTracking),
      uppercaseTextPct: pct(upperChars, chars),
      rowGapMax: round1(rowGaps.reduce((a, b) => Math.max(a, b), 0)), rowGapMedian: round1(median(rowGaps)),
      rowGapLow: round1(quantile(rowGaps, 0.1)), rowGapHigh: round1(quantile(rowGaps, 0.9)),
      blockGapMax: round1(blockGaps.reduce((a, b) => Math.max(a, b), 0)), blockGapMedian: round1(median(blockGaps)),
      blockGapLow: round1(quantile(blockGaps, 0.1)), blockGapHigh: round1(quantile(blockGaps, 0.9)),
      contentWidth: Math.round(contentWidth), contentWidthPct: pct(contentWidth, viewportWidth),
      paletteColours: colours.size,
      transitions, animations,
      layoutTablePct: share(layout, 'table'), layoutFloatPct: share(layout, 'float'), layoutFlexPct: share(layout, 'flex'),
      layoutGridPct: share(layout, 'grid'), layoutAbsolutePct: share(layout, 'absolute'),
      textDensity: docW * docH ? round1(chars / (docW * docH / 1000)) : 0,
      mediaAreaPct: pct(mediaArea, docW * docH),
    },
  });
    return result;
  } finally {
    if (arrayToJSON) Array.prototype.toJSON = arrayToJSON;
    for (const view of views) view.style.removeProperty('display');
  }
})()