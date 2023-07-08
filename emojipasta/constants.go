package emojipasta

import (
	_ "embed"
	"encoding/json"
	// "unicode/utf8"
	// "github.com/enescakir/emoji"
)

//go:embed mappings.json
var emojiData []byte
var mappings map[string][]string
// var emojiMappings map[string]string

func init() {
	json.Unmarshal(emojiData, &mappings)
	// emojiMappings = make(map[string]string, len(mappings))
	// for k, v := range mappings {
	// 	currentData := []byte{}
	// 	for _, char := range v {
	// 		currentData = utf8.AppendRune(currentData, char)
	// 	}
	// 	emojiMappings[k] = string(currentData)
	// }

}