package parser

import "github.com/decayofmind/zsh-fast-alias-tips/pkg/model"

func Parse(line string) model.AliasDef {
	alias := make([]rune, 0, 1024)
	expanded := make([]rune, 0, 1024)

	afterEscape := false
	inQuote := false
	inRightExp := false
	for _, aRune := range line {
		if aRune == '\\' {
			afterEscape = !afterEscape

			if afterEscape {
				continue
			}
		}

		if aRune == '\'' && !afterEscape {
			inQuote = !inQuote
		} else if aRune == '=' && !inQuote {
			inRightExp = true
		} else if !inRightExp {
			alias = append(alias, aRune)
		} else {
			expanded = append(expanded, aRune)
		}

		afterEscape = false
	}

	return model.AliasDef{Name: string(alias), Expanded: string(expanded)}
}
