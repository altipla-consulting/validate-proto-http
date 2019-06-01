#!/bin/bash

set -eu

REPO_NAME=validate-proto-http

URL=$(curl -sSL https://api.github.com/repos/altipla-consulting/$REPO_NAME/releases/latest | jq -r '.assets[].browser_download_url')
echo "Downloading latest release from: $URL"
curl $URL -sSL -o /usr/local/bin/$REPO_NAME || (
  echo ""
  echo " [!!!] This command needs sudo priviledge"
  echo ""
  exit 1
)

chmod +x,o+x /usr/local/bin/$REPO_NAME

echo ""
echo "Run the app with:"
echo "    $REPO_NAME"
echo ""
