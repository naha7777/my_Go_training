package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	runes := []rune(strings.TrimSpace(string(data)))

	elem := strings.Fields(string(runes))

	count := 0
	mi := math.MaxInt
	ma := math.MinInt

	for i, e := range elem {
		if i != 0 {
			res, err := strconv.Atoi(string(e))
			if err != nil {
				fmt.Println(err)
				return
			}
			count += res
			if res < mi {
				mi = res
			}
			if res > ma {
				ma = res
			}
		}
	}
	fmt.Printf("somme=%d min=%d max=%d", int(count), int(mi), int(ma))
}
