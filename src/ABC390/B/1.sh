#!/usr/bin/env bash
set -eEuo pipefail
# set -xv


readarray -t stdin
readarray -t nums < <(tr " " "\n" <<<"${stdin[1]}" | tac)

index=0
prehi=0
for n in "${nums[@]}"; do
    if (( index == "${#nums[@]}" - 1 )); then
        break
    fi
    current="$n"

    next="${nums["$((index+1))"]}"

    # hi=$(( current / next ))
    hi=$(bc <<< "scale=10; $current / $next")
    if [[ "$prehi" != 0 ]]  && [[ "$prehi" != "$hi" ]]; then
        echo "No"
        exit 0
    fi
    prehi="$hi"
    index=$((index + 1))
done
echo Yes
