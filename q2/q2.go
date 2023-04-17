package main

import (
	"fmt"
	"strings"
)

func main() {

	AverageLettersPerWord("O rato roeu a roupa do rei de Roma")

}

func AverageLettersPerWord(text string) (float64, error) {

	palavras := strings.Split(text, " ")
	soma := 0.
	for x := 0; x < len(palavras); x++ {
		soma += float64(len(palavras[x]))
	}

	media := soma / float64(len(palavras))

	fmt.Printf("%f", media)

	return 0, nil
}
