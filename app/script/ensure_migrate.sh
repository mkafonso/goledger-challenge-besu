#!/usr/bin/env bash

set -e

if ! command -v migrate &> /dev/null; then
  echo "❌ migrate not found. Installing..."

  OS="$(uname -s)"
  ARCH="$(uname -m)"

  if [ "$OS" = "Darwin" ]; then
    echo "macOS detected"

    if ! command -v brew &> /dev/null; then
      echo "Homebrew not found. Please install brew first."
      exit 1
    fi

    brew install golang-migrate

  else
    echo "Linux detected"
    curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xvz
    sudo mv migrate /usr/local/bin/
  fi

else
  echo "✅ migrate already installed"
fi