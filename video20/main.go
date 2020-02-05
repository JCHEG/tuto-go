package main

import (
	"fmt"
)

func main() {
	printInfoNoParam()
	printInfoAvecParam("Jonny", "begood@gmail.com", 50)
	result := moyenne(15.5, 14.5)
	fmt.Printf("La moyenne est %f\n", result)

	somme := sommeNameReturned(10, 12, 5)
	fmt.Printf("La somme est %d\n", somme)
}

func printInfoNoParam() {
	fmt.Printf("Nom => %s , email => %s, age => %d\n", "Jamel", "jamel@gamil.com", 44)
}

func printInfoAvecParam(nom string, email string, age int) {
	fmt.Printf("Nom => %s , email => %s, age => %d\n", nom, email, age)
}

// on peut typer plusieurs variables en même temps
func moyenne(x, y float64) float64 {
	return (x + y) / 2
}

// On peut aussi initialiser le nom de la vairable de retour dans la signature de la fonction
func sommeNameReturned(x, y, z int) (somme int) {
	somme = x + y + z
	return somme
}
