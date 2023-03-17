package grammar

import (
	"regexp"
	"strconv"
	"unicode"
)

var ordinalNumberDictionary = map[int]string{
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

func CapitalizeFirstLetter(sentence string) string {
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
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return num + ordinalNumberDictionary[n%10]
}

func OrdinalizeNumbers(sentence string) string {
	return regexp.MustCompile("[0-9]+").ReplaceAllStringFunc(sentence, convertToOrdinal)
}

func SentenceTokenizer(text string) []string {
	return regexp.MustCompile("[^.!?]*[.!?]").FindAllString(text, -1)
}
