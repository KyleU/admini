#!/bin/bash

## Tags the git repo

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

TGT=$1

[ "$TGT" ] || (echo "must provide one argument, like \"0.0.1\"" && exit)

make build

git add .
git commit -m "v${TGT}" || true

git tag "v${TGT}"

git push
git push --tags
