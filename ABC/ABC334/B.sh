#!/bin/bash

# なんかしらんけどACできません！！！

readarray -t stdin < <(tr " " "\n")
base="${stdin[0]}"
perM="${stdin[1]}"
Tkhs="${stdin[2]}"
Aoki="${stdin[3]}"

function calc() {
    awk " BEGIN{OFMT=\"%.22f\"; print $*}"
}

function getInt() {
    cut -d "." -f 1 <<<"$1"
}

function getLeftK() {
    local n
    n=$(calc "($1 - $base) / $perM")

    # 正のとき切り捨て
    if [[ $(getInt "$n") != -* ]]; then
        n=$(getInt "$n")
    else
        # 負のとき切り上げ
        n=-$((-$(getInt "$n") + 1))
    fi
    echo "$n"

}

function getRightK() {
    echo $(($(getLeftK "$1") + 1))
}

output=$((-"$(getRightK "$Tkhs")" + "$(getLeftK "$Aoki")" + 1))

echo $output
