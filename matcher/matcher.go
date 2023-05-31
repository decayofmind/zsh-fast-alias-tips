package matcher

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
)

func Match(defs []model.AliasDef, command string) *model.AliasDef {
	sort.Slice(defs, func(i, j int) bool {
		return len(defs[j].Expanded) <= len(defs[i].Expanded)
	})

	var candidate model.AliasDef

	for {
		var match model.AliasDef
		for _, def := range defs {

			if command == def.Expanded || strings.HasPrefix(command, def.Expanded) {
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
		return &candidate
	}

	return nil
}
