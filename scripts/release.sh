#!/bin/bash

set -e

RELEASE="v$1"
CURRENT="$(git tag | tail -n1)"

echo "Installing needed tools..."
go get github.com/aktau/github-release

echo "Creating Release"
mkdir -p release
go build podfetcher.go
mv podfetcher release
cp scripts/Makefile release
cp examples/config.toml release
cp examples/feeds release

mv release podfetcher

tar cfav podfetcher.tar.bz2 podfetcher
rm -rf podfetcher

LOG="$(git log --pretty=oneline --abbrev-commit "$CURRENT"..HEAD)"
git tag "$RELEASE"
git push origin "$RELEASE"
github-release release \
--user gregf \
--repo podfetcher \
--tag "$RELEASE" \
--name "podfetcher" \
--description "$LOG" \

github-release upload \
--user gregf \
--repo podfetcher \
--tag "$RELEASE" \
--name "podfetcher.tar.bz2" \
--file "podfetcher.tar.bz2"

rm -f podfetcher.tar.bz2
