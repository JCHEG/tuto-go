package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("*****************************************")

	// Boucle while
	j := 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("*****************************************")
	// Boucle forever
	k := 0
	for {
		k++
		if k%2 != 0 {
			fmt.Printf("k est impair et vaut %d\n", k)
			continue
		}
		if k > 10 {
			fmt.Printf("k est pair et supérieur à 10 et vaut %d\n", k)
			break
		}
	}
}
