#!/bin/bash

set -eu

URL=$(curl -sSL https://api.github.com/repos/altipla-consulting/validate-proto-http/releases/latest | jq -r '.assets[].browser_download_url')
echo "Downloading latest release from: $URL"
curl $URL -sS -o /usr/local/bin/validate-proto-http || (
  echo ""
  echo " [!!!] This command needs sudo priviledge"
  echo ""
  exit 1
)
