#!/usr/bin/env bash

readarray -t stdin < <(tr " " "\n")

n="${stdin[0]}"
x="${stdin[1]}"
array=()
for al in {A..Z}; do
    while read -r; do
        readarray -t -O "${#array[@]}" array <<<"$al"
    done < <(seq 1 "$n")
done
{
    printf "%s" "${array[@]}"
    echo
} | fold -w 1 | head -n "$x" | tail -n 1
