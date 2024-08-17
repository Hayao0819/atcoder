#!/usr/bin/env bash

readarray -t stdin < <(tr " " "\n")
total="${stdin[0]}"
takahasi="${stdin[1]}"
aoki="${stdin[2]}"
unchecked=$(( total - takahasi - aoki ))

if (( takahasi == aoki )); then
    echo No
elif (( takahasi > aoki )); then
    if (( takahasi > aoki + unchecked )); then
        echo Yes
    else
        echo No
    fi
elif (( takahasi < aoki )); then
    if (( takahasi + unchecked < aoki )); then
        echo Yes
    else
        echo No
    fi
fi
