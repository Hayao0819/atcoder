#!/usr/bin/env bash

set -eEuo pipefail
declare -A dlist=(["N"]="S" ["S"]="N" ["E"]="W" ["W"]="E")

while read -r d; do
    printf "%s" "${dlist[$d]}"
done < <(fold -w 1) && echo
