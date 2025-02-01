#!/usr/bin/env python3

import sys

def main():
    try:
        stdin = sys.stdin.read().strip().split("\n")
        nums = list(map(float, stdin[1].split()))[::-1]
        prehi = 0
        if 0 in nums:
            print("No")
        
        for index in range(len(nums) - 1):
            current = nums[index]
            next_num = nums[index + 1]
            hi = current / next_num

            if prehi != 0 and prehi != hi:
                print("No")
                return

            prehi = hi

        print("Yes")

    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()
