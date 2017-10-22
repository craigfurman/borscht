#!/bin/bash
set -euo pipefail

(
cd "$(dirname "$0")/.."
ginkgo -r -randomizeAllSpecs -randomizeSuites "$@"
)
