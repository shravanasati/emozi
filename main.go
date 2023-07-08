package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shravan-1908/emozi/emojipasta"
)

const helpText = `
emozi is a simple command line tool to insert emojis in between text. it can read input
from stdin as well as from the given arguments.

visit "https://github.com/Shravan-1908/emozi".. for more information.
`

func main() {
	gen := emojipasta.New().WithDefaultMappings()

	// Check if there is any input available in stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			input := scanner.Text()
			fmt.Println(gen.GenerateEmojiPasta(input))
		}

		if err := scanner.Err(); err != nil {
			return
		}
	} else {
		if len(os.Args) > 1 {
			args := os.Args[1:]
			text := strings.Join(args, " ")
			fmt.Println(gen.GenerateEmojiPasta(text))
		} else {
			fmt.Println(gen.GenerateEmojiPasta(helpText))
		}
	}
}
