#!/usr/bin/env bash
# shellcheck source=../git/core.bash # @root

_ferret=$(which ferret || true)

# Example: run ferret exec fetch.fql | jq
ferret() {
  [ -z "${_ferret}" ] && @fatal Please setup environment first

  pushd "$(@root)/tests/requests" >/dev/null || exit 1
  trap 'popd >/dev/null' ERR

  "${_ferret}" "$@"

  popd >/dev/null || exit 1
}
