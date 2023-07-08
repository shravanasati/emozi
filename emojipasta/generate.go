package emojipasta

import (
	"math/rand"
	"strings"
	"time"
)

const maxEmojisPerBlock = 2

type Generator struct {
	mappings map[string][]rune
}

func New() *Generator {
	return &Generator{}
}

func (epg *Generator) WithDefaultMappings() *Generator {
	epg.mappings = emojiMappings
	return epg
}

func (epg *Generator) WithCustomMappings(customMapping map[string][]string) *Generator {
	epg.mappings = processMapping(customMapping)
	return epg
}

func (epg *Generator) GeneratePasta(text string) string {
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

	var final strings.Builder
	for _, v := range newBlocks {
		final.WriteString(v)
	}
	return final.String()
}

func (epg *Generator) generateEmojis(block string) []rune {
	trimmedBlock := trimNonalphabeticalCharacters(block)
	if itemInSlice(trimmedBlock, commonWords) {
		return []rune{}
	}
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	matchingEmojis := epg.getMatchingEmojis(trimmedBlock)
	emojis := []rune{}

	if len(matchingEmojis) > 0 {
		numEmojis := gen.Intn(maxEmojisPerBlock) + 1
		for i := 0; i < numEmojis; i++ {
			choice := matchingEmojis[gen.Intn(len(matchingEmojis))]
			emojis = append(emojis, choice)
		}
	}
	return emojis
}

func (epg *Generator) getMatchingEmojis(trimmedBlock string) []rune {
	key := getAlphanumericPrefix(strings.ToLower(trimmedBlock))
	_, ok := epg.mappings[key]
	if ok {
		return epg.mappings[getAlphanumericPrefix(key)]
	}
	return []rune{}
}

func getAlphanumericPrefix(s string) string {
	i := 0
	for i < len(s) && isAlnum(rune(s[i])) {
		i++
	}
	return s[:i]
}

func isAlnum(r rune) bool {
	return ('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9')
}
