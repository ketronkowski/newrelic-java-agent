#!/usr/bin/env bash

set -euo pipefail

env GOOS="linux" go build -ldflags='-s -w' -o bin/helper github.com/paketo-buildpacks/newrelic-java-agent/cmd/helper
env GOOS="linux" go build -ldflags='-s -w' -o bin/main github.com/paketo-buildpacks/newrelic-java-agent/cmd/main

if [ "${STRIP:-false}" != "false" ]; then
  strip bin/helper bin/main
fi

if [ "${COMPRESS:-none}" != "none" ]; then
  $COMPRESS bin/helper bin/main
fi

ln -fs main bin/build
ln -fs main bin/detect
