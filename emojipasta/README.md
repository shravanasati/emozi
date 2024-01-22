# emojipasta

emojipasta is a minimal library that can generate emojipasta from the given text.

This is the library which powers [emozi](https://github.com/shravanasati/emozi).

[![Go Reference](https://pkg.go.dev/badge/github.com/shravanasati/emozi/emojipasta.svg)](https://pkg.go.dev/github.com/shravanasati/emozi/emojipasta)


## Installation

```
go get github.com/shravanasati/emozi/emojipasta
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/shravanasati/emozi/emojipasta"
)

func main() {
	generator := emojipasta.New().WithDefaultMappings()
	// you can also set your custom mappings
	// generator := emojipasta.New().WithCustomMappings(myCustomMapping)

	// the default max emojis per block is set to 2
	// you can change it as following
	err := generator.SetMaxEmojisPerBlock(3)
	if err != nil {
		...
	}
	fmt.Println(generator.GenerateEmojiPasta("I just hope this works. No tests are available sorry."))
}
```
