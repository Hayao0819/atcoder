#!/usr/bin/env bash

readarray -t line < <(tr " " "\n")

TakoyakiTime="${line[0]}"
GoBedTime="${line[1]}"
GetUpTime="${line[2]}"

echo "${line[@]}" >&2

sleepingTimes=()

# 8時に寝て、10時に起きる
if (( GoBedTime < GetUpTime )); then
    readarray -t sleepingTimes < <(seq "$GoBedTime" "$GetUpTime")
else
    readarray -t sleepingTimes < <(seq "$GoBedTime" 24 && seq 0 "$GetUpTime")
fi

if printf "%s\n" "${sleepingTimes[@]}" | grep -qx "$TakoyakiTime"; then
    echo No
else
    echo Yes
fi
