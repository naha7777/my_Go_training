package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
		return
	}

	str := strings.TrimSpace(string(data))
	count := 1
	dic := make(map[string]int)

	for _, c := range str {
		if unicode.IsSpace(c) == false {
			val, ok := dic[string(c)]
			if ok == true {
				dic[string(c)] = val + 1
			} else {
				dic[string(c)] = count
			}
		}
	}

	s := fmt.Sprint(dic)
	s = strings.TrimPrefix(s, "map[")
	s = strings.TrimSuffix(s, "]")
	s = strings.ReplaceAll(s, " ", "\n")
	s = strings.ReplaceAll(s, ":", ": ")

	fmt.Print(s)
}
