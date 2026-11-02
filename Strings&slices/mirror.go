package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := strings.TrimSpace(string(data))

	// Convertir la chaîne en un slice de runes (caractères Unicode)
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
}
