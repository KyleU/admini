#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

time gon ./tools/notarize/gon.amd64.hcl
