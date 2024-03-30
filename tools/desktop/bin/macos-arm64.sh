#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Uses `tools/desktop` to build a desktop application, intended to be run from Docker

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/.."

echo "starting macOS arm64 desktop build..."
GOOS=darwin GOARCH=arm64 CC=aarch64-apple-darwin23-clang CXX=aarch64-apple-darwin23-clang++ go build -o ../../dist/darwin_arm64/admini
