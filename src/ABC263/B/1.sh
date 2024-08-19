#!/usr/bin/env bash
set -Eeuo pipefail
readarray -t stdin

ninzu="${stdin[0]}"
readarray -t parents < <(tr " " "\n" <<<"${stdin[1]}")

count=1
find_parent() {
    local p="${parents[$(($1 - 2))]}"
    if ((p == 1)); then
        return 0
    fi
    ((count++))
    find_parent "$p"
}

find_parent "$ninzu"

echo "$count"
