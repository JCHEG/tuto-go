package main

import (
	"fmt"
)

func main() {

	var noms [3]string
	fmt.Printf("Noms => %v, taille du tableau noms => %v\n", noms, len(noms))

	noms[0] = "Jonny"
	noms[1] = "Diego"
	noms[2] = "Ronaldo"
	fmt.Printf("Noms => %v, taille du tableau noms => %v\n", noms, len(noms))

	//Une autre manière de déclarer des tableaux
	impairs := [4]int{1, 3, 5, 7}
	fmt.Printf("Tableau d'impaire => %v, taille du tableau impairs => %v\n", impairs, len(impairs))

}
