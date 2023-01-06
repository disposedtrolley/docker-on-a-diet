package main

import (
	"fmt"
	"log"
	"os"

	git "github.com/libgit2/git2go/v31"
)

func main() {
	repoDir := os.Args[1]
	repo, err := git.OpenRepository(repoDir)
	if err != nil {
		log.Fatal(err)
	}

	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	commit, err := repo.LookupCommit(head.Target())
	if err != nil {
		log.Fatal(err)
	}

	for commit != nil {
		fmt.Printf("[%s <%s>] @ %s\n", commit.Author().Name, commit.Author().Email, commit.Author().When)
		fmt.Println(commit.Message())

		commit = commit.Parent(0)
	}
}
