package main

import (
	"fmt"
	"io"
	"os"
	"math"
	"strconv"
	"strings"
)

func square(nb float64) float64 {
	res := math.Pow(nb, 2)
	return res
}

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

	ch := make(chan float64)
	var cal float64

	go func() {
		for i := 1; i <= n; i++ {
			res := square(float64(i))
			cal += res
		}
		ch <- cal
	}()

	val := <-ch

	fmt.Println(int(val))
}

// afficher la somme des carres de 1 a N
