#!/usr/bin/env bash

readarray -t stdin < <(tr " " "\n")
sum=$(("${stdin[0]}" + "${stdin[1]}"))
if ((sum >= 10)); then
    echo error
else
    echo "$sum"
fi
