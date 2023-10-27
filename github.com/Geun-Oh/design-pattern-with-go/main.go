package main

import (
	"fmt"

	"github.com/Geun-Oh/design-pattern-with-go/singleton"
)

func main() {

	for i := 0; i < 20; i++ {
		go singleton.GetInstance()
	}

	fmt.Scanln()

	return
}
