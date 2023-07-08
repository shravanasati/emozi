package emojipasta

import (
	"regexp"
)

var (
	blockRegex = regexp.MustCompile(`\s*[^\s]*`)
	trimRegex  = regexp.MustCompile(`^\W*|\W*$`)
)

// A 'block' is a prefix of whitespace characters followed
// by a series of non-whitespace characters.
func splitIntoBlocks(text string) []string {
	if text == "" || blockRegex.FindStringIndex(text) == nil {
		return []string{text}
	}
	blocks := []string{}
	start := 0
	for start < len(text) {
		blockMatch := blockRegex.FindStringIndex(text[start:])
		block := text[start : start+blockMatch[1]]
		blocks = append(blocks, block)
		start += blockMatch[1]
	}
	return blocks
}

func trimNonalphabeticalCharacters(text string) string {
	return trimRegex.ReplaceAllString(text, "")
}

func itemInSlice[T comparable](item T, slice []T) bool {
	for _, val := range slice {
		if item == val {
			return true
		}
	}
	return false
}