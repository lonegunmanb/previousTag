package main

import (
	"fmt"
	"github.com/lonegunmanb/previousTag/pkg"
	"os"
)

func main() {
	owner := os.Args[1]
	repo := os.Args[2]
	currentTag := os.Args[3]

	tag, err := pkg.PreviousTag(owner, repo, currentTag)
	if err != nil {
		panic(err)
	}
	fmt.Print(tag)
}
