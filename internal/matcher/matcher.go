package matcher

import (
	"fmt"
	"sort"
	"strings"

	"github.com/decayofmind/zsh-fast-alias-tips/internal/model"
)

func Match(defs []model.AliasDef, command string) (*model.AliasDef, bool) {
	sort.Slice(defs, func(i, j int) bool {
		return len(defs[j].Expanded) <= len(defs[i].Expanded)
	})

	var candidate model.AliasDef
	isFullMatch := false

	for {
		var match model.AliasDef
		for _, def := range defs {

			if command == def.Expanded {
				match = def
				isFullMatch = true
				break
			} else if strings.HasPrefix(command, def.Expanded) {
				match = def
				break
			}
		}

		if match != (model.AliasDef{}) {
			command = fmt.Sprintf("%s%s", match.Name, command[len(match.Expanded):])
			candidate = match
			continue
		}

		break
	}

	if candidate != (model.AliasDef{}) {
		return &candidate, isFullMatch
	}

	return nil, isFullMatch
}
