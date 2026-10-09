#!/bin/sh
# Checks one token variant against its era's pack: builds a throwaway
# library, composes the era's structure with the variant's palette, type
# and surface, lints the result, checks the Never rules against its own
# specimen and writes a screenshot.
# usage: scripts/variant-check.sh <era-slug> <nn> [out.png]   (nn as in packs/<era>/tokens/<nn>)
set -eu
era=$1
nn=$2
shot=${3:-${SHOTS:-shots}/$era-variant-$nn.png}
[ -d "packs/$era/tokens/$nn" ] || { echo "no packs/$era/tokens/$nn" >&2; exit 2; }

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
mkdir -p "$work/src/$era/tokens"
cp "packs/$era/era.json" "$work/src/$era/"
cp -R "packs/$era/01" "$work/src/$era/01"
cp -R "packs/$era/tokens/$nn" "$work/src/$era/tokens/02"
bin/pack build-library -version check -out "$work/lib" "$work/src" >/dev/null

id=$(sed -n 's/.*"id": *"\([^"]*\)".*/\1/p' "packs/$era/01/style.json" | sed 's/p01\.t01\.u01$/p02.t02.u02/')
bin/pack compose -library "$work/lib" -out "$work/pack" "$id"
status=0
bin/pack lint "$work/pack" | sed "s|$work/pack|variant $nn|" || status=1
if bin/pack adhere -pack "$work/pack" "$work/pack/specimen.html" | grep "unusable rule"; then status=1; fi
width=$(sed -n 's/.*"width": *\([0-9]*\).*/\1/p' "$work/pack/style.json")
mkdir -p "$(dirname "$shot")"
bin/pack shot -width "$width" -height 2400 "$work/pack/specimen.html" "$shot"
echo "screenshot: $shot"
exit $status
