package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Println(err)
		return
	}

	if n % 2 == 0 {
		fmt.Print("pair")
	} else {
		fmt.Print("impair")
	}
}
