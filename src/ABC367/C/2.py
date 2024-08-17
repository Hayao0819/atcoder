#!/usr/bin/env python3

import sys

def sum_digits(num_str):
    return sum(int(x) for x in num_str.replace(" ", ""))

def makelist(fixed, limits, len_limits):
    space_removed = fixed.replace(" ", "")
    fixed_clm = len(space_removed)
    
    if fixed_clm >= len_limits:
        return
    
    result = []
    for i in range(1, limits[fixed_clm] + 1):
        v = f"{fixed} {i}".strip()
        if fixed_clm + 1 == len_limits:
            result.append(v)
        else:
            result.extend(makelist(v, limits, len_limits))
    
    return result

def main():
    stdin = sys.stdin.read().splitlines()
    times = int(stdin[0].split()[1])
    limits = list(map(int, stdin[1].split()))
    len_limits = len(limits)
    
    candidates = makelist("", limits, len_limits)
    answer = [num for num in candidates if sum_digits(num) % times == 0]
    
    for ans in answer:
        print(ans)

if __name__ == "__main__":
    main()
