#!/usr/bin/env python3

import sys


def parse_score(l: list[str]) -> list:
    return [int(l[0]), l[1]]


stdin = sys.stdin.readlines()
n = int(stdin[0])
score = [parse_score(x.rsplit("\n")[0].split(" ")) for x in stdin[1:]]
# print(score)


def get_moves_by_hand(hand: str) -> list[int]:
    return [x[0] for x in score if x[1] == hand]


def main() -> None:
    rmoves = get_moves_by_hand("R")
    lmoves = get_moves_by_hand("L")

    tired = 0
    for c in range(len(rmoves) - 1):
        tired = tired + abs(rmoves[c] - rmoves[c + 1])
    for c in range(len(lmoves) - 1):
        tired = tired + abs(lmoves[c] - lmoves[c + 1])
    print(tired)


if __name__ == "__main__":
    main()
