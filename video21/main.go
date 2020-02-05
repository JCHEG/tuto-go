package main

import (
	"fmt"
	"strings"
)

func main() {
	nom, longueur := passerEnMiniscule("RONALDO")
	fmt.Printf("La mot %s a été passé en minuscule et il comporte %d lettres\n", nom, longueur)

	// On peut ignorer un (ou des valeurs retourné)
	_, longueur2 := passerEnMiniscule("MARADONNA")
	fmt.Printf("La mot a été passé en minuscule et il comporte %d lettres\n", longueur2)

	nom2, _ := passerEnMiniscule("ZINEDINE ZIDANE")
	fmt.Printf("La mot %s a été passé en minuscule et il comporte des lettres\n", nom2)
}

func passerEnMiniscule(mot string) (string, int) {
	return strings.ToLower(mot), len(mot)
}
