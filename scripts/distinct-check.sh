#!/bin/sh
# Reports token sets of an era that are too alike: two palettes, two type
# sets or two surface sets that differ in fewer than a third of their
# tokens. Mixing parts is only worth offering when each choice shows.
# usage: scripts/distinct-check.sh <era-slug>
set -eu
era=$1
set --
for f in "packs/$era"/[0-9][0-9]/tokens.css "packs/$era"/tokens/*/tokens.css; do
	[ -f "$f" ] && set -- "$@" "$f"
done
awk '
FNR == 1 { set = FILENAME; sub(/\/tokens\.css$/, "", set); sets[++n] = set; part = "" }
/@part/ { part = $0; sub(/.*@part +/, "", part); sub(/[ *\/]+$/, "", part); next }
# --figure-shift is an optical correction, not part of the character of a set.
part != "" && /^[ \t]*--[a-z0-9-]+[ \t]*:/ && !/--figure-shift/ {
	line = $0; sub(/^[ \t]*/, "", line); name = line; sub(/[ \t]*:.*/, "", name)
	value = line; sub(/^[^:]*:[ \t]*/, "", value); sub(/;.*$/, "", value)
	v[part, set, name] = value; names[part, name] = 1; parts[part] = 1
}
END {
	bad = 0
	for (p in parts) {
		total = 0
		for (k in names) { split(k, a, SUBSEP); if (a[1] == p) total++ }
		for (i = 1; i <= n; i++) for (j = i + 1; j <= n; j++) {
			d = 0
			for (k in names) { split(k, a, SUBSEP); if (a[1] == p && v[p, sets[i], a[2]] != v[p, sets[j], a[2]]) d++ }
			if (d * 3 < total) { printf "%s: %s and %s differ in %d of %d tokens\n", p, sets[i], sets[j], d, total; bad = 1 }
		}
	}
	if (!bad) print "ok: every palette, type set and surface set of " ERA " is distinct"
	exit bad
}' ERA="$era" "$@"
