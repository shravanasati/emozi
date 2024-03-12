package emojipasta

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

var ErrNegativeEmojisPerBlock = errors.New("cannot set negative emojis per block")

// Generator holds the emoji mappings and the number of max emojis per block.
type Generator struct {
	mappings          map[string][]rune
	maxEmojisPerBlock int
}

// GeneratorBuilder is used to build the Generator.
type GeneratorBuilder struct {
	generator *Generator
	err       error
}

// Returns a pointer to a Generator struct with maxEmojisPerBlock set to 2.
func NewBuilder() *GeneratorBuilder {
	// using the builder pattern here
	// if the error is not nil, all subsequent calls to building methods
	// will return early, while ensuring safe chaining
	return &GeneratorBuilder{
		err: nil,
		generator: &Generator{maxEmojisPerBlock: 2},
	}
}

// Sets Generator.mappings to the default emoji mappings.
func (builder *GeneratorBuilder) WithDefaultMappings() *GeneratorBuilder {
	if builder.err != nil {
		return builder
	}
	builder.generator.mappings = emojiMappings
	return builder
}

// Sets Generator.mappings to the given custom mappings.
// Example of a custom mapping is {"hi": ["✋", "👋"], "person": ["👦", "🧔"]...}
func (builder *GeneratorBuilder) WithCustomMappings(customMapping map[string][]string) *GeneratorBuilder {
	if builder.err != nil {
		return builder
	}
	builder.generator.mappings = processMapping(customMapping)
	return builder
}

// Sets Generator.maxEmojisPerBlock to the given number. Call to the build method will return an
// error is the given number is negative.
func (builder *GeneratorBuilder) WithMaxEmojisPerBlock(n int) *GeneratorBuilder {
	if builder.err != nil {
		return builder
	}
	if n < 0 {
		builder.err = ErrNegativeEmojisPerBlock
		return builder
	}
	builder.generator.maxEmojisPerBlock = n
	return builder
}

// Returns the generator if no errors were encountered during the building process (adding options).
func (builder *GeneratorBuilder) Build() (*Generator, error) {
	if builder.err != nil {
		return nil, builder.err
	}
	return builder.generator, nil
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
