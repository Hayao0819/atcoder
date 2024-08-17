#!/usr/bin/env bash

readarray -t stdin
times=$(cut -d " " -f 2 <<< "${stdin[0]}")
readarray -t limits < <(tr " " "\n" <<< "${stdin[1]}")
len="${#limits[@]}"

# sum 123 -> 6
sum(){
    #shellcheck disable=SC2005,SC1102
    echo "$(( $(printf "%s+" "${@}" ) + 0 ))"
}


makelist(){
    local fixed="$1"
    local spaceRemoved="${fixed// /}"
    local fixedClm="${#spaceRemoved}"
    if (( fixedClm >= len));then
        return
    fi
    local i v
    for i in $(seq 1 "${limits["$fixedClm"]}" ); do
        v="$fixed $i"
        if (( fixedClm + 1 == len )); then
            echo "$v"
        fi
        makelist "$v"
    done
}

answer=()
while read -r num; do
    if  (( "$( eval "sum $num"  )" % times == 0 )); then
        answer+=("$num")
    fi
done < <(makelist "" | sed "s|^ ||g") 
printf "%s\n" "${answer[@]}"
