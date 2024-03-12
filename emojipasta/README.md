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
	generator, err := emojipasta.NewBuilder().
		WithDefaultMappings().
		WithMaxEmojisPerBlock(3). // default max emojis per block is 2
		Build()

	if err != nil {
		...
	}

	// you can also set your custom mappings
	// generator, err := emojipasta.NewBuilder().WithCustomMappings(myCustomMapping)

	fmt.Println(generator.GenerateEmojiPasta("I just hope this works. No tests are available sorry."))
}
```
