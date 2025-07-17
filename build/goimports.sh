#!/bin/sh

find_files() {
  find . ! \( \
      \( \
        -path '.github' \
      \) -prune \
    \) -name '*.go' \
    -not -name '*_gen.go'
}

GOFMT="gofmt -s -w"
GOIMPORTS="goimports -w"

find_files | xargs sed -i '' '/^import (/,/^)/{/^$/d;}'
find_files | xargs $GOFMT
find_files | xargs $GOIMPORTS
