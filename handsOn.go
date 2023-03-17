package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func capitalizeFirstLetter(sentence string) string {
	mutableSentence := []rune(sentence)
	for index, r := range mutableSentence {
		if unicode.IsLetter(r) {
			mutableSentence[index] = unicode.ToUpper(r)
			break
		}
	}
	return string(mutableSentence)
}

func convertToOrdinal(num string) string {
	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return num + ordinalDictionary[n%10]
}

func ordinalizeNumbers(sentence string) string {
	return regexp.MustCompile("[0-9]+").ReplaceAllStringFunc(sentence, convertToOrdinal)
}

func sentenceTokenizer(text string) []string {
	return regexp.MustCompile("[^.!?]*[.!?]").FindAllString(text, -1)
}

func main() {
	text := "hello to my 1 attempt. this is Me!  give me the code.are you Señor? ? you bought 22, 23 and 25 of them."
	sentences := sentenceTokenizer(text)
	result := ""
	for _, sentence := range sentences {
		sentence = ordinalizeNumbers(sentence)
		result += capitalizeFirstLetter(sentence)
	}
	fmt.Print(result)
}
