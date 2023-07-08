package emojipasta

import (
	"fmt"
	"unicode/utf8"
)

func PrintMappings() {
	for k, v := range mappings {
		ems := []rune{}
		for _, val := range v {
			em, _ := utf8.DecodeRuneInString(val)
			ems = append(ems, em)
		}
		fmt.Printf("%s %c \n", k, ems)
	}
}