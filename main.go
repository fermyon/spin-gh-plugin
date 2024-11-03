package main

import (
	"log"

	"github.com/fermyon/spin-gh-plugin/cmd/gh"
)

func init() {
	log.SetFlags(0)
}

func main() {
	gh.ExecuteRootCommand()
}
