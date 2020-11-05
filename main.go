package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	numWords, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing word count, please supply an integer")
		os.Exit(1)
	}

	rolls := rollDice(numWords)
	fmt.Println(translateRolls(rolls))
}

func rollDice(c int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var rollSlice []string
	for i := 0; i < c; i++ {
		var rollString string
		for i := 0; i <= 5; i++ {
			rollString = rollString + fmt.Sprint(r.Intn(5)+1)
		}
		rollSlice = append(rollSlice, rollString)
	}
	return rollSlice
}

func translateRolls(s []string) string {
	var diceString string
	dictionary, err := ioutil.ReadFile("source_dictionary")
	if err != nil {
		fmt.Println("Error while reading dictionary file")
		os.Exit(1)
	}

	return diceString
}
