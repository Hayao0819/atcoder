#!/usr/bin/env bash
cnt=0
readarray -t input < <(tail -n 1 | tr " " "\n")
while true; do
    for ((i = 0; i < ${#input[@]} - 1; i++)); do
        if (("${input[$i]}" % 2 == 0)); then
            input[${i}]=$(("${input[$i]}" / 2))

        else
            break 2
        fi
    done
    cnt=$((cnt + 1))
done
echo "${cnt}"
