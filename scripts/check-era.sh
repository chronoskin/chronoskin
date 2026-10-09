#!/bin/sh
# Everything a layout or token set must pass before it is proposed, for one
# era: the linter and the Never rules of each pack, a click through the
# views of each specimen in a browser, token sets that differ
# from one another, and every layout composed with every token set.
# Screenshots of the compositions are written to shots/ (or the folder
# named by SHOTS) to look at.
# With a layout number, only that layout is composed (the whole era takes
# several minutes).
# usage: scripts/check-era.sh <era-slug> [layout-number]
set -eu
era=${1:?usage: scripts/check-era.sh <era-slug> [layout-number]}
only=${2:-}
[ -x bin/pack ] || make build

status=0
for pack in "packs/$era"/[0-9][0-9]; do
	[ -d "$pack" ] || continue
	bin/pack lint "$pack" || status=1
	bin/pack adhere -pack "$pack" "$pack/specimen.html" || status=1
done
bin/pack views "packs/$era"/[0-9][0-9] || status=1
sh scripts/distinct-check.sh "$era" || status=1
sh scripts/structure-check.sh "$era" $only || status=1

if [ $status -eq 0 ]; then echo "$era: ok"; else echo "$era: FAILED" >&2; fi
exit $status
