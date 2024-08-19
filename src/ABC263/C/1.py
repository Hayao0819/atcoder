#!/usr/bin/env python3
import sys

stdin: list[str] = sys.stdin.readlines()[0].split(" ")
length = int(stdin[0])
max = int(stdin[1])


def makelist(fixed: list[int]) -> None:
    for i in range(1, max + 1):
        if len(fixed) == length:
            print(" ".join([str(x) for _, x in enumerate(fixed)]))
            return
        if len(fixed) > 0 and fixed[len(fixed) - 1] >= i:
            continue
        new = fixed + [i]
        makelist(new)


if __name__ == "__main__":
    makelist([])
