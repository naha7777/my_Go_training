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

	str := strings.TrimSpace(string(data))
	words := strings.Fields(str)
}

// un stockage par personne
// faut que je stocke le nom de la personne
// la somme de ses notes
// puis combien de notes elle a eu
// pour ensuite faire la somme de ses notes / combien de notes elle a eu
