package main

import (
	"distributedComputing-CA1/grammar"
	"fmt"
)

func main() {
	text := "hello to my 1 attempt. this is Me!  give me the code.are you Señor? ? you bought 22, 23 and 25 of them."
	sentences := grammar.SentenceTokenizer(text)
	result := ""
	for _, sentence := range sentences {
		sentence = grammar.OrdinalizeNumbers(sentence)
		result += grammar.CapitalizeFirstLetter(sentence)
	}
	fmt.Print(result)
}
