package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"strconv"
	"strings"
)

type Task struct {
	Index	int
	N		int
}

func factorielle(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func main() {
	// extraire la data de l'entrée
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
		return
	}

	// découper la data
	elem := strings.Fields(string(data))
	if len(elem) == 0 {
		return
	}

	// convertir en entier chaque element de la data sauf le 1er
	nums := make([]int, 0, len(elem)-1)
	for _, s := range elem[1:] {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Print(err)
			return
		}
		nums = append(nums, v)
	}

	// créer un channel de taches
	tasks := make(chan Task)
	// créer un dictionnaire pour stocker l'index et le resultat
	res := make([]int, len(nums))
	// créer le compteur de goroutines
	var wg sync.WaitGroup

	for w := 0; w < 4; w++ {
		// ajouter un worker
		wg.Add(1)
		// lancer la fonction dans un go routine séparée
		go func() {
			// signale que le worker a fini en fin de fonction
			defer wg.Done()
			// pour chaque tache on stock son index et le resultat de la
			// la factorielle de N dans le dictionnaire de resultats
			for t := range tasks {
				res[t.Index] = factorielle(t.N)
			}
		}()
	}

	// rempli le channel de task
	for i, v := range nums {
		tasks <- Task{Index: i, N: v}
	}
	// ferme le channel après l'envoi de toutes les taches
	close(tasks)

	// bloque jusqu'a ce que le compteur soit a 0 (donc attend que les workers)
	// finissent leur taff
	wg.Wait()

	// print final
	for i, v := range nums {
		fmt.Printf("%d! = %d\n", v, res[i])
	}
}
