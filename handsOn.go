package main

import (
	"bufio"
	"distributedComputing-CA1/grammar"
	"fmt"
	"log"
	"os"
)

func main() {
	//text := "hello to my 1 attempt. this is Me!  give me the code.are you Señor? ? you bought 22, 23 and 25 of them."
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	sentences := grammar.SentenceTokenizer(scanner.Text())
	result := ""
	for _, sentence := range sentences {
		sentence = grammar.OrdinalizeNumbers(sentence)
		result += grammar.CapitalizeFirstLetter(sentence)
	}
	fmt.Print(result)
}
