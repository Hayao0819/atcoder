#!/usr/bin/env bash
readarray -t stdin
readarray -t head < <(printf "%s\n" "${stdin[@]}" | head -n 1 | tail -n 1 | tr " " "\n")
readarray -t goals < <(printf "%s\n" "${stdin[@]}" | head -n 2 | tail -n 1 | tr " " "\n")

foodCount=${head[0]}
#nutritionCount=${head[1]}

readarray -t foodLines < <(printf "%s\n" "${stdin[@]}" | tail -n "$foodCount" | tr " " "-")

#echo "${head[@]}" >&2
#echo "${goals[@]}" >&2
#echo "${foodLines[@]}" >&2

nutritionSum=()
for foodLine in "${foodLines[@]}"; do
    readarray -t parsed < <(echo "$foodLine" | tr "-" "\n")
    clmCount=0
    for food in "${parsed[@]}"; do
        nutritionSum["$clmCount"]=$(( ${nutritionSum[$clmCount]} + "$food" ))
        clmCount=$(( clmCount + 1 ))
    done
done

#set -xv

clmCount=0
for goal in "${goals[@]}"; do
    clmSum="${nutritionSum["$clmCount"]}"
    if (( goal > clmSum )); then
        echo "No"
        exit 0
    fi
    clmCount=$(( clmCount + 1 ))
done
echo "Yes"
