package main

import (
	"fmt"
	"github.com/libgit2/git2go/v34"
	"log"
)

func main() {

	repo, err := git.OpenRepository(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repo)
}
