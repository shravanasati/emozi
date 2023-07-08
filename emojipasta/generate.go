package emojipasta

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Generator holds the emoji mappings and the number of max emojis per block.
type Generator struct {
	mappings          map[string][]rune
	maxEmojisPerBlock int
}

// Returns a pointer to a Generator struct with maxEmojisPerBlock set to 2.
func New() *Generator {
	return &Generator{maxEmojisPerBlock: 2}
}

// Sets Generator.mappings to the default emoji mappings.
func (epg *Generator) WithDefaultMappings() *Generator {
	epg.mappings = emojiMappings
	return epg
}

// Sets Generator.mappings to the given custom mappings.
// Example of a custom mapping is {"hi": ["✋", "👋"], "person": ["👦", "🧔"]...}
func (epg *Generator) WithCustomMappings(customMapping map[string][]string) *Generator {
	epg.mappings = processMapping(customMapping)
	return epg
}

// Sets Generator.maxEmojisPerBlock to the given number. Returns an error is the given
// number is negative.
func (epg *Generator) SetMaxEmojisPerBlock(n int) error {
	if n < 0 {
		return (fmt.Errorf("cannot set negative max emoji per block: %d", n))
	}
	epg.maxEmojisPerBlock = n
	return nil
}

// Generates the emoji pasta from the given text.
func (epg *Generator) GenerateEmojiPasta(text string) string {
	blocks := splitIntoBlocks(text)
	newBlocks := []string{}
	for _, block := range blocks {
		var b strings.Builder
		b.WriteString(block)
		emojis := epg.generateEmojis(block)
		for _, em := range emojis {
			b.WriteRune(em)
		}
		newBlocks = append(newBlocks, b.String())
	}

	return strings.Join(newBlocks, "")
}

// returns emojis for the given block.
func (epg *Generator) generateEmojis(block string) []rune {
	trimmedBlock := trimNonalphabeticalCharacters(block)
	if itemInSlice(trimmedBlock, commonWords) {
		return []rune{}
	}
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	matchingEmojis := epg.getMatchingEmojis(trimmedBlock)
	emojis := []rune{}

	if len(matchingEmojis) > 0 {
		numEmojis := randGen.Intn(epg.maxEmojisPerBlock) + 1
		if numEmojis > len(matchingEmojis) {
			numEmojis = len(matchingEmojis) - 1
		}
		for i := 0; i < numEmojis; i++ {
			choice := matchingEmojis[randGen.Intn(len(matchingEmojis))]
			emojis = append(emojis, choice)
		}
	}
	return emojis
}

// returns all matching emojis for the given block.
func (epg *Generator) getMatchingEmojis(trimmedBlock string) []rune {
	key := getAlphanumericPrefix(strings.ToLower(trimmedBlock))
	_, ok := epg.mappings[key]
	if ok {
		return epg.mappings[getAlphanumericPrefix(key)]
	}
	return []rune{}
}

// returns the alphanumeric prefix from the given string.
func getAlphanumericPrefix(s string) string {
	i := 0
	for i < len(s) && isAlnum(rune(s[i])) {
		i++
	}
	return s[:i]
}

// checks if the given rune is either a letter or a number.
func isAlnum(r rune) bool {
	return ('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9')
}
