#!/usr/bin/env python3

import sys

stdin = sys.stdin.readlines()
n = int(stdin[0])
arr = [int(x) for x in stdin[1].split(" ")]


def is_tousa(l: list[int]):
    return ((l[0] + l[len(l) - 1]) * len(l)) / 2 == sum(l)


# lists: list[list[int]] = []
count=0

def generate_all_pattern(fixed: list[int]):
    for i in range(1, n + 1):
        if len(fixed) == 2:
            l, r = fixed[0]-1, fixed[1]
            if is_tousa( arr[l:r]):
                global count 
                count+=1
            return
        if len(fixed) > 0 and fixed[len(fixed) - 1] > i:
            continue
        new = fixed + [i]
        generate_all_pattern(new)

generate_all_pattern([])
print(count)



