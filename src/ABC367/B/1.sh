#!/usr/bin/env bash

readarray -t X < <(tr "." "\n")

seisu="${X[0]}"
shosu="${X[1]}"

while true; do
    zeroRemoved=${shosu%0}
    [[ "$zeroRemoved" = "$shosu" ]] && break
    shosu="$zeroRemoved"
done

if [[ -z "$shosu" ]]; then
    echo "$seisu"
else
    echo "${seisu}.${shosu}"
fi

