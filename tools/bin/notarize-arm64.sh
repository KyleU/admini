#!/bin/bash

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || (echo "must provide one argument, like \"0.0.1\"" && exit)

if [ "$XSKIP_EXTRA" != "true" ]
then
  if [ "$XSKIP_NOTARIZE" != "true" ]
  then
    sed -i '' "s/v[01]\.[0-9]*[0-9]\.[0-9]*[0-9][-SNAPSHOT]*/$TGT/g" ./tools/notarize/gon.arm64.hcl
    time gon ./tools/notarize/gon.arm64.hcl
    sed -i '' "s/v[01]\.[0-9]*[0-9]\.[0-9]*[0-9][-SNAPSHOT]*/v0.0.0/g" ./tools/notarize/gon.arm64.hcl
  fi
fi
