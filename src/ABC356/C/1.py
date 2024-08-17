#!/usr/bin/env python3

import sys
import os


# stdin to array
stdin: list[str] = sys.stdin.readlines()
keynum:int = int(stdin[0].split()[0])
testnum:int = int(stdin[0].split()[1])
correctnum:int = int(stdin[0].split()[2])
rawtestcases:list[list] = [x.split() for x in stdin[1:]]

# 2進数の01をboolのリストで表現
def intToBoolArray(num: int) -> list[bool]:
    ret = []
    bnum = bin(num)
    dig = len(bnum)
    for i in range(0, len(bnum)-2):
        ret.append((bnum[i+2]) == '1')
    return ret

# Arrayの長さが一緒になるようにFalseを追加
def addFalse(arr: list[bool], num: int) -> list[bool]:
    ret = arr
    for i in range(0, num-len(arr)):
        ret = [False] + ret
    return ret

# 配列の各要素の最大の長さを返す
def maxArrayLength(arr: list[list]) -> int:
    ret = 0
    for i in range(0, len(arr)):
        if len(arr[i]) > ret:
            ret = len(arr[i])
    return ret

def getAllPattern(keynum: int) -> list[list[bool]]:
    allpattern = []
    for i in range (0, 2**keynum):
        allpattern.append(intToBoolArray(i))

    maxlen = maxArrayLength(allpattern)
    for i in range(0, len(allpattern)):
        allpattern[i] = addFalse(allpattern[i], maxlen)

    return allpattern

def main():
    allpattern = getAllPattern(keynum)
    print(allpattern)

main()
