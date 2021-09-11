#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

## Starts the app, reloading on changes

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

[[ -f "$HOME/bin/oauth" ]] && . $HOME/bin/oauth
export admini_encryption_key=TEMP_SECRET_KEY

ulimit -n 2048
air
