package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	numWords, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing word count, please supply an integer")
		os.Exit(1)
	}

	rolls := rollDice(numWords)
	// translateRolls(rolls)
	fmt.Println(translateRolls(rolls))
}

func rollDice(c int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var rollSlice []string
	for i := 0; i < c; i++ {
		var rollString string
		for i := 0; i <= 4; i++ {
			rollString = rollString + fmt.Sprint(r.Intn(5)+1)
		}
		rollSlice = append(rollSlice, rollString)
	}
	return rollSlice
}

func translateRolls(s []string) string {
	var diceString string
	for i := range s {
		// start a go routine to find each word
		// go findWord()

		diceString = diceString + findWord(s[i])
	}
	return diceString
	// listen on a channel to append the results to diceString

}

func findWord(s string) string {
	dictionary, err := os.Open("source_dictionary")
	if err != nil {
		fmt.Println("Error while reading dictionary file")
		os.Exit(1)
	}
	defer dictionary.Close()
	var word string
	var foundWord = false

	scanner := bufio.NewScanner(dictionary)
	for scanner.Scan() {
		if foundWord == true {
			break
		}
		if strings.Contains(scanner.Text(), s) {
			foundWord = true
			word = scanner.Text()[6:] + " "
		}
	}
	return word
}
