#!/usr/bin/env python3

import sys

stdin = sys.stdin.read().split(" ")
a = int(stdin[0])
b = int(stdin[1])


# def is_tousa(l: list[int]):
#     return l[1] - l[0] == l[2] - l[1]


def abx() -> int:
    return b - (b - a)


def axb() -> float:
    return (a + b) / 2


def bax() -> int:
    return a - (a - b)


def eq() -> bool:
    return a == b


def main() -> None:
    # a b x
    # a x b
    # b a x

    xlist: list[int] = []

    for x, _ in enumerate([abx(), bax()]):
        if not x in xlist:
            xlist.append(x)

    leng = len(xlist)

    if axb().is_integer() and not eq():
        leng = leng + 1
    if eq():
        leng = 1
    # print(xlist)
    print(leng)


if __name__ == "__main__":
    main()
