package main

import (
	"fmt"
)

func main() {
	//jour dans un mois
	jour := 26
	switch jour {
	case 10:
		fmt.Println("Nous sommes le 10")
	case 18:
		fmt.Println("Nous sommes le 18")
	default:
		fmt.Printf("Nous sommes un autre jour, nous sommes le %d \n", jour)
	}

	// On peut faire un switch sans valeur de dépar à tester
	// Les tests se feront dans les cases
	heure := 19
	fmt.Printf("Il est %d heure\n", heure)
	switch {
	case heure < 12:
		fmt.Println("Nous sommes le matin")
	case heure > 12 && heure < 18:
		fmt.Println("Nous sommes l'après midi")
	default:
		fmt.Println("Nous sommes le soir")
	}
}
