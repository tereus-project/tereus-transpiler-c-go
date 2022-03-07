package main

import (
	"fmt"
	"os"

	"github.com/tereus-project/tereus-remixer-c-go/remixer"
)

func main() {
	if len(os.Args) >= 2 {
		output, err := remixer.Remix(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(output)
	}
}
