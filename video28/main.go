package main

import (
	"fmt"
)

func main() {

	start()
	// finish()
	// fait éxécuter finish à la fin (à la sortie de la fonction)
	defer finish() //LIFO: Last In First Out
	names := []string{"Chirstiano", "Ronaldo", "Diego", "Zizou"}
	for _, name := range names {
		fmt.Printf("Salut %v !\n", name)
		defer fmt.Printf("Au revoir %v !\n", name)
	}

}

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}
