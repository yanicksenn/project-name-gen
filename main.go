package main

import (
	"embed"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//go:embed data/adjectives.txt
//go:embed data/nouns.txt
var content embed.FS

func main() {
	adjectives := ReadAdjectives()
	nouns := ReadNouns()

	for i := 0; i < GetAmountToGenerate(); i++ {
		randomAdjective := GetRandomValueFrom(adjectives)
		randomNoun := GetRandomValueFrom(nouns)
		fmt.Println(fmt.Sprintf("%s %s", randomAdjective, randomNoun))
	}
}

func ReadAdjectives() []string {
	return ReadTextFile("data/adjectives.txt")
}

func ReadNouns() []string {
	return ReadTextFile("data/nouns.txt")
}

func ReadTextFile(filePath string) []string {
	text, _ := content.ReadFile(filePath)
	valuesRaw := strings.Split(string(text), "\n")
	values := []string{}

	for i := 0; i < len(valuesRaw); i++ {
		var value = valuesRaw[i]
		if !strings.HasPrefix(value, "#") {
			values = append(values, value)
		}
	}

	return values
}

func GetAmountToGenerate() int {
	if len(os.Args) > 1 {
		amountToGenerate, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(errors.New("amount to generate: Must be a positive integer"))
		}
		if amountToGenerate < 1 {
			panic(errors.New("amount to generate: Must not be less than 1"))
		}

		return amountToGenerate
	}

	return /* default amount to generate = */ 1
}

func GetRandomValueFrom(array []string) string {
	return array[GetRandomValueBetween(0, len(array)-1)]
}

func GetRandomValueBetween(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
