#!/bin/bash
set -euo pipefail

if [ $# -ne 1 ]; then
  echo "run with image"
  exit
fi

IMAGE=$1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ctr=$(buildah from scratch)
mnt=$(buildah mount "${ctr}")

cp markadr "$mnt"/

buildah config --entrypoint "[\"/markadr\"]" "${ctr}"
buildah config -p 8123 "${ctr}"
buildah config -e ASSET_PATH=/srv/app "${ctr}"

buildah commit "${ctr}" "${IMAGE}"
buildah unmount "${ctr}"
