#!/bin/sh
# Checks every layout of an era with every token set of that era: builds a
# throwaway library from the era's folder (packs/<era>), composes each structure with each
# palette/type/surface set, lints the result and writes a screenshot. For a
# structure with its own tokens it also checks the Never rules.
# With NOSHOTS=1 no screenshots are taken, which is much quicker; otherwise
# they go to shots/, or to the folder named by SHOTS.
# usage: [NOSHOTS=1] scripts/structure-check.sh <era-slug> [structure-number]
set -eu
era=$1
only=${2:-}
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
shots=${SHOTS:-shots}
mkdir -p "$work/src" "$shots"
cp -R "packs/$era" "$work/src/"
bin/pack build-library -version check -out "$work/lib" "$work/src"

structures=$(ls "$work/lib/eras/$era/structures")
sets=$(ls "$work/lib/eras/$era/palettes" | sed 's/\.json$//')
status=0
for s in $structures; do
	[ -n "$only" ] && [ "$s" != "$only" ] && continue
	# Layouts of one era may be different kinds of page: the page type in
	# the ID is the one of this layout's own pack, the s-th in order.
	archetype=$(sed -n 's/.*"archetype": *"\([^"]*\)".*/\1/p' "packs/$era/$s/style.json")
	for k in $sets; do
		id="v1.$era.$archetype.s$s.p$k.t$k.u$k"
		out="$work/pack-$s-$k"
		bin/pack compose -library "$work/lib" -out "$out" "$id"
		if ! bin/pack lint "$out" | sed "s|$out|layout $s, tokens $k|"; then status=1; fi
		if [ "$s" = "$k" ]; then
			if bin/pack adhere -pack "$out" "$out/specimen.html" | grep "unusable rule"; then status=1; fi
		fi
		[ -n "${NOSHOTS:-}" ] && continue
		width=$(sed -n 's/.*"width": *\([0-9]*\).*/\1/p' "$out/style.json")
		# One screenshot per view of the specimen.
		for view in $(sed -n 's/.*<section class="ds-view[^"]*" id="\([a-z0-9-]*\)".*/\1/p' "$out/specimen.html"); do
			bin/pack shot -width "$width" -height 2600 "file://$out/specimen.html#$view" "$shots/$era-s$s-k$k-$view.png"
			echo "screenshot: $shots/$era-s$s-k$k-$view.png"
		done
	done
done
exit $status
