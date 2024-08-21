package main

import (
	"fmt"
	"io"

	// "log/slog"
	"os"
	"strconv"
	"strings"
)

func handle(err error) {
	if err != nil {
		//fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func sumArr(nums ...int) int {
	r := 0
	for _, i := range nums {
		r = r + i
	}
	return r
}

// slices.Contains
func contains[T comparable](l []T, i T) bool {
	for _, e := range l {
		if e == i {
			return true
		}
	}
	return false
}

func deleteFunc[T any](l []T, f func(i T) bool) []T {
	narr := []T{}
	for _, e := range l {
		if !f(e) {
			narr = append(narr, e)
		}
	}
	return narr
}

var count int = 0

func sumCheck(nums []int) {
	//slog.Info("sumCheck", "nums", nums, "count", count, "len", len(nums))

	if len(nums) == 0 {
		return
	}

	sum := sumArr(nums...)
	sum3mod := sum % 3
	if sum3mod == 0 {
		fmt.Println(count)
		os.Exit(0)
	}

	mods := []int{sum3mod}
	for {
		last := mods[len(mods)-1]
		if last >= sum {
			break
		}
		mods = append(mods, last+3)
	}
	for _, i := range mods {
		//slog.Info("Checking", "i", i)
		if contains(nums, i) {
			isRemoved := false
			removed := deleteFunc(nums, func(j int) bool {
				if isRemoved {
					return false
				}
				if i == j {
					isRemoved = true
					return true
				}
				return false
			})
			count = count + 1
			sumCheck(removed)
		}
	}
}

func timesCheck(nums []int) {
	times := []int{}
	for _, e := range nums {
		if e%3 == 0 {
			times = append(times, e)
		}
	}
	if len(times) != 0 {
		fmt.Println(len(nums) - len(times))
		os.Exit(0)
	}
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	handle(err)

	nums := []int{}
	for _, r := range string(bytes) {
		s := strings.TrimSpace(string(r))
		if s == "" {
			break
		}

		n, err := strconv.Atoi(s)
		handle((err))
		nums = append(nums, n)
	}

	sumCheck(nums)

	timesCheck(nums)

	fmt.Println(-1)
	os.Exit(0)
}
