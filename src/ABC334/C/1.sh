#!/bin/bash
#shellcheck disable=SC2016

readarray -t stdin
readarray -t lost < <(tr " " "\n" <<<"${stdin[1]}" | xargs -I{} bash -c 'echo $(( {} - 1 ))')

printf "%s\n" "${lost[@]}"
