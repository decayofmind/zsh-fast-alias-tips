package parser

import (
	"testing"

	"github.com/decayofmind/zsh-fast-alias-tips/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestParse_NoQuotesInRightExp(t *testing.T) {
	assert.Equal(t, Parse("dk=docker"), model.AliasDef{Name: "dk", Expanded: "docker"})
}

func TestParse_QuotesInRightExp(t *testing.T) {
	assert.Equal(t, Parse("gb='git branch'"), model.AliasDef{Name: "gb", Expanded: "git branch"})
}

func TestParse_QuotesInBothExps(t *testing.T) {
	assert.Equal(t, Parse("'g cb'='git checkout -b'"), model.AliasDef{Name: "g cb", Expanded: "git checkout -b"})
}
