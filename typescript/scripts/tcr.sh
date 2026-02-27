#!/usr/bin/env bash
# TCR — Test && Commit || Revert
# Si los tests pasan, se hace commit automático.
# Si los tests fallan, se revierten todos los cambios no commiteados.
# Uso: ./tcr.sh [path]

set -o pipefail

TEST_PATH="${1:-}"

if [ -n "$TEST_PATH" ]; then
  echo "🧪 Running tests in: $TEST_PATH"
  npx vitest run "$TEST_PATH" --reporter=default 2>&1
else
  echo "🧪 Running tests..."
  npx vitest run --reporter=default 2>&1
fi

if [ $? -eq 0 ]; then
  echo "✅ Tests passed — committing changes."
  git add -A && git commit -m "TCR: green"
else
  echo "❌ Tests failed — reverting changes."
  git checkout .
  git clean -fd
fi
