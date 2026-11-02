package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := io.ReadAll((os.Stdin))
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(int(i))
		}
	}
}
