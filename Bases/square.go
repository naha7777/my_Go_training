package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
		return
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Print(err)
		return
	}

	res := n*n

	fmt.Print(res)
}
