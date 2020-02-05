package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	occ, lines, text, err := findReplaceFile("wikigo.txt", "Go", "Python")

	fmt.Println(occ)
	fmt.Println(lines)
	fmt.Println(err)
	fmt.Println(text)

}

func findReplaceFile(src, old, new string) (occ int, lines []int, text string, err error) {
	dat, errorFile := os.Open(src)
	defer dat.Close()
	if errorFile != nil {
		fmt.Printf("Message d'erreur => %v\n", errorFile)
		err = errorFile
	}
	numLigne := 0
	// var text string = ""
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		numLigne++
		ligne := scanner.Text()
		found, res, occTemp := processLine(ligne, old, new)
		if found {
			occ += occTemp
			lines = append(lines, numLigne)
			text += res + "\n"
		}
	}
	return occ, lines, text, err
}

func processLine(line, old, new string) (found bool, res string, occ int) {
	if strings.Contains(line, old) {
		mots := strings.Fields(line)
		for i, mot := range mots {
			if strings.EqualFold(mot, old) {
				re, _ := regexp.Compile(`[A-Z]`)
				if re.MatchString(string(mot[0])) {
					mots[i] = strings.Title(new)
				} else {
					mots[i] = strings.ToLower(new)
				}
				occ++
				found = true
			}
		}
		res = strings.Join(mots, " ")
		return found, res, occ
	}
	return found, res, occ
}
