package main

import "fmt"

func main() {
	// LES SLICES
	// On crée le slice
	nums := make([]int, 2, 3)

	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v => len = %v et cap = %v \n", nums, len(nums), cap(nums))

	nums = append(nums, 64)
	fmt.Printf("%v => len = %v et cap = %v \n", nums, len(nums), cap(nums))

	// La capacité double dés que la taille du tableau est depassée
	nums = append(nums, 42)
	fmt.Printf("%v => len = %v et cap = %v \n", nums, len(nums), cap(nums))

	nums = append(nums, 25)
	nums = append(nums, 36)
	nums = append(nums, 22)
	nums = append(nums, 127)
	fmt.Printf("%v => len = %v et cap = %v \n", nums, len(nums), cap(nums))
	fmt.Println("*****************************************")

	//En ne précisant pas la taille du tableau c'est un slice qui est créé de manière implicite
	langage := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v => len = %v et cap = %v \n", langage, len(langage), cap(langage))

	langage = append(langage, "!")
	fmt.Printf("%v => len = %v et cap = %v \n", langage, len(langage), cap(langage))
	fmt.Println("*****************************************")

	//!!!!!!!!!!!! Si on crée de sub slice ils pointeront vers le même tableau que celui créé par le slice parent
	subSlice1 := langage[:2]
	subSlice2 := langage[3:]
	fmt.Printf("%v => len = %v et cap = %v \n", langage, len(langage), cap(langage))
	fmt.Printf("%v => len = %v et cap = %v \n", subSlice1, len(subSlice1), cap(subSlice1))
	fmt.Printf("%v => len = %v et cap = %v \n", subSlice2, len(subSlice2), cap(subSlice2))
	fmt.Println("*****************************************")

	langage[0] = "UP"
	langage[6] = "?"
	fmt.Printf("%v => len = %v et cap = %v \n", langage, len(langage), cap(langage))
	fmt.Printf("%v => len = %v et cap = %v \n", subSlice1, len(subSlice1), cap(subSlice1))
	fmt.Printf("%v => len = %v et cap = %v \n", subSlice2, len(subSlice2), cap(subSlice2))
	fmt.Println("*****************************************")
	//Clonage d'un slice
	sliceCopy := make([]string, len(langage))
	copy(sliceCopy, langage)
	sliceCopy[0] = "DOWN"
	fmt.Printf("%v => len = %v et cap = %v \n", langage, len(langage), cap(langage))
	fmt.Printf("%v => len = %v et cap = %v \n", sliceCopy, len(sliceCopy), cap(sliceCopy))

}
