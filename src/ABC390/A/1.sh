#!/usr/bin/env bash
set -eEuo pipefail
num="$(tr -d " ")"
currect=(21345 13245 12435 12354)
if printf "%s\n" "${currect[@]}" | grep -qx "$num"; then
    echo "Yes"
else
    echo "No"
fi
