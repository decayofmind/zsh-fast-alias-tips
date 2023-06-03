package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/decayofmind/zsh-fast-alias-tips/pkg/matcher"
	"github.com/decayofmind/zsh-fast-alias-tips/pkg/model"
	"github.com/decayofmind/zsh-fast-alias-tips/pkg/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		os.Exit(1)
	}

	aliases := make([]model.AliasDef, 0, 1024)
	excludes := strings.Split(os.Getenv("ZSH_FAST_ALIAS_TIPS_EXCLUDES"), " ")

	scanner := bufio.NewScanner(bufio.NewReaderSize(os.Stdin, 1024))
	for scanner.Scan() {
		alias := parser.Parse(scanner.Text())

		skip := false

		for _, exclude := range excludes {
			if exclude == alias.Name {
				skip = true
			}
		}

		if !skip {
			aliases = append(aliases, alias)
		}
	}

	command := os.Args[1]
	match := matcher.Match(aliases, command)
	if match != nil {
		fmt.Printf("%s%s\n", match.Name, command[len(match.Expanded):])
	}
}
