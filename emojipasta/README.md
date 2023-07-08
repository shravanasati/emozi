# emojipasta

emojipasta is a minimal library that can generate emojipasta from the given text.

This is the library which powers [emozi](https://github.com/Shravan-1908/emozi).

## Installation

```
go get github.com/Shravan-1908/emozi/emojipasta
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Shravan-1908/emozi/emojipasta"
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
	generator.GenerateEmojiPasta("I just hope this works. No tests are available sorry.")
}
```
