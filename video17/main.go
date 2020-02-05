package main

import (
	"fmt"
)

func main() {
	//jour dans un mois
	jour := 25
	if jour < 15 {
		fmt.Printf("Nous sommes dans la prmeière quinzaine du mois (jour=%d)\n", jour)
	} else if jour == 18 {
		fmt.Printf("Nous sommes dans un jour spécliale, le %d \n", jour)
	} else {
		fmt.Printf("Nous sommes juste le %d \n", jour)
	}

	// declaration et intitlaisation d'un variable dans un if
	// cette variable n'est accessible que depuis le bloc du if
	if count := 10; count > 5 {
		fmt.Println("Count est superieur à 5. Count vaut :", count)
	}
}
