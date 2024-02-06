#!/usr/bin/env python3
import sys
import re

def main():
    r = re.compile(r"^0+|0+$")
    num=sys.stdin.readline().replace("\n","")
    num=r.sub("",num)
    for i, n in enumerate(num):
        #print(f"{i}:{len(num)-1-i}")
        if num[i] != num[len(num)-1-i]:
            print("No")
            return
    print("Yes")

if __name__ == '__main__':
    main()
