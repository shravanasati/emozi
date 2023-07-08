package main

import (
	"fmt"

	"github.com/Shravan-1908/emozi/emojipasta"
)

func main()  {
	fmt.Println("lmao ded")
	generator := emojipasta.New().WithDefaultMappings()
	fmt.Println(generator.GeneratePasta("harry potter is just a weakling, he knows nothing but angry"))
}