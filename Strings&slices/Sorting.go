package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"slices"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	elem := strings.Fields(string(data))

	nums := make([]int, 0, len(elem))
	for i, f := range elem {
		if i != 0 {
			n, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			nums = append(nums, n)
		}
	}

	slices.Sort(nums)

	for i, n := range nums {
		fmt.Print(n)
		if i != len(elem) - 1 {
			fmt.Print(" ")
		}
	}
}
