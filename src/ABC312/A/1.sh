#!/bin/bash
target="$(cat)"
words=(ACE BDF CEG DFA EGB FAC GBD)
if printf "%s\n" "${words[@]}" | grep -q "$target"; then
    echo "Yes"
else
    echo "No"
fi
