package emojipasta

import (
	_ "embed"
	"encoding/json"
	"unicode/utf8"
)

//go:embed mappings.json
var emojiData []byte
var mappings map[string][]string

var emojiMappings map[string][]rune

var commonWords = []string{
	"a",
	"an",
	"as",
	"is",
	"if",
	"of",
	"the",
	"it",
	"its",
	"or",
	"are",
	"this",
	"with",
	"so",
	"to",
	"at",
	"was",
	"and",
}

func processMapping(givenMapping map[string][]string) map[string][]rune {
	newMapping := make(map[string][]rune, len(givenMapping))

	for k, v := range mappings {
		ems := []rune{}
		for _, val := range v {
			em, _ := utf8.DecodeRuneInString(val)
			ems = append(ems, em)
		}
		newMapping[k] = ems
	}
	return newMapping
}

func init() {
	json.Unmarshal(emojiData, &mappings)
	emojiMappings = processMapping(mappings)
}
