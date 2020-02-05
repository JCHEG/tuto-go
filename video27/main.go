package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dat, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("Message d'erreur => %v\n", err)
		return
	}

	fmt.Printf(dat)

}

func readFile(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	if len(dat) == 0 {
		// return "", errors.New("Erreur personnalisée => Fichier vide")
		//Erreur avec du contenu formaté
		return "", fmt.Errorf("Erreur personnalisée => Fichier %v vide", fileName)
	}

	return string(dat), nil
}
