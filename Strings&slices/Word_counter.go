package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := strings.TrimSpace(string(data))

	runes := []rune(str)

	count := 0

	for i := 0; i <= len(runes) - 1; i++ {
		if unicode.IsLetter(runes[i]) && i + 1 < len(runes) && unicode.IsSpace(runes[i+1]) {
			count += 1
		} else if unicode.IsLetter((runes[i])) && i == len(runes) - 1 {
			count += 1
		}
	}

	fmt.Print(count)
}
