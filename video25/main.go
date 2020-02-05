package main

import (
	"fmt"
)

func main() {
	names := []string{"Jonny", "Diego", "Christiano", "Romario", "Mohamed"}
	for index, value := range names {
		fmt.Printf("Nous sommes dans le tour de boucle numéro %d et le nom courant est %s\n", index, value)
	}
	//pour ignorer l'index on peut le remplacer par "_"
	for _, lettre := range "Chaine de caractère" {
		fmt.Printf("La lettre courante est %v\n", string(lettre))
	}
}
