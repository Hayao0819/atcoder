#!/usr/bin/env python3

import sys

def getIndex(arr: list, index: int):
    tlist = []
    for i, v in enumerate(arr):
        if (v):
            tlist.append(i)
    return tlist[index]

def main():
    # 設定読み取り
    stdin = sys.stdin.read().splitlines()
    gridscount = int(stdin[0].split(" ")[0])

    # Grid初期化
    grid = [False]*gridscount

    # Gridに初期値を設定
    for a in map(lambda x: int(x)-1, stdin[1].split(" ")):
        grid[a]=True

    # 操作をQ回適用
    for l in map(lambda x:int(x)-1, stdin[2].split(" ")):
        # l番目のコマのあるマス目
        index =getIndex(grid, l)

        if gridscount == index+1:
            continue

        if (grid[index+1] == False):
            grid[index]=False
            grid[index+1]=True

    out = []
    for i,g in enumerate(grid):

        if g == True:
            out.append(str(i+1))
        
    print(" ".join(out))
            
if __name__ == "__main__":
    main()
