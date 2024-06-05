#!/usr/bin/env bash

readarray -t stdin
readarray -t goal < <( printf "%s\n" "${stdin[@]}" | head -n 1 | tr " " "\n")
echo "${goal[@]}"
