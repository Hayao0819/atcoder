#!/usr/bin/env bash

readarray -t input < <(tr " " "\n")

len="${input[0]}"
revStart="${input[1]}"
revEnd="${input[2]}"

readarray -t ret < <(seq 1 "$len")

for i in $(seq 1 "$len"); do
    if ((i >= revStart)) && ((i <= revEnd)); then
        ret["$(("$i" - 1))"]="$((revEnd - i + revStart))"
    fi
done

echo "${ret[@]}"
