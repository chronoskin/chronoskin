#!/bin/sh
# Starts a new layout for an era: copies the era's first pack to the next
# free number, so that you begin from a pack that already passes the linter
# and follows the era's conventions. Then rewrite it: a new kind of page,
# and a palette, type set and surface set of its own.
# usage: scripts/new-layout.sh <era-slug>
set -eu
era=${1:?usage: scripts/new-layout.sh <era-slug>}
first="packs/$era/01"
[ -d "$first" ] || { echo "no era $era in packs" >&2; exit 1; }

# Token sets without a layout are numbered after the packs, so the new
# layout takes the first number that neither has.
n=2
while [ -d "packs/$era/$(printf %02d "$n")" ] || [ -d "packs/$era/tokens/$(printf %02d "$n")" ]; do n=$((n + 1)); done
nn=$(printf %02d "$n")
new="packs/$era/$nn"

cp -R "$first" "$new"
# The Style ID names the pack's four parts; they all take the new number.
for f in "$new"/*; do
	sed "s/\.s01\.p01\.t01\.u01/.s$nn.p$nn.t$nn.u$nn/g" "$f" > "$f.tmp" && mv "$f.tmp" "$f"
done

echo "$new: a copy of $first, as layout $nn"
echo "next: rewrite its five files, then run scripts/check-era.sh $era $nn"
