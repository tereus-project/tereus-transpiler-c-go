package main

import (
	"os"

	"github.com/tereus-project/tereus-remixer-c-go/remixer"
)

func main() {
	if len(os.Args) >= 2 {
		e := remixer.Remix(os.Args[1])
		if e != nil {
			panic(e)
		}
	}
}
