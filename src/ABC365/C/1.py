#!/usr/bin/env python

import sys

stdin: list[str] = sys.stdin.readlines()
max = int(stdin[0].split(" ")[1])
money = [int(x) for x in stdin[1].split(" ")]

def calc(m: int) -> int:
    return sum(min([x, m]) for _, x in enumerate(money))


# 予算オーバーならTrue
def isMaxOver(m: int) -> bool:
    r = calc(m)
    return r > max


# めぐる式二分探索
# 0|--------|------|--------|-----|2*10**14
#         left→   Goal   ←right
# Goal: 高橋くんの予算
# left: 予算内
# right: 予算外
def main() -> None:
    right: int = 2 * 10**14
    left: int = 1

    moneysum = sum(money)
    if max >= moneysum:
        print("infinite")
        exit()

    while right - left > 1:
        mid = int((right + left) / 2)

        if isMaxOver(mid):
            right = mid
        else:
            left = mid

    print(left)


if __name__ == "__main__":
    main()
