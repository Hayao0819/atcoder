#!/usr/bin/env bash

readarray -t stdin < <(tr " " "\n")
(("${stdin[0]}" > "${stdin[1]}")) && {
    echo Bat
    exit 0
}
echo Glove
