package matcher

import (
	"testing"

	"github.com/decayofmind/zsh-fast-alias-tips/pkg/model"
	"github.com/stretchr/testify/assert"
)

var mockAliasDefs = []model.AliasDef{
	{
		Name:     "dk",
		Expanded: "docker",
	},
	{
		Name:     "gb",
		Expanded: "git branch",
	},
	{
		Name:     "gco",
		Expanded: "git checkout",
	},
	{
		Name:     "gcb",
		Expanded: "git checkout -b",
	},
	{
		Name:     "ls",
		Expanded: "ls -G",
	},
	{
		Name:     "ll",
		Expanded: "ls -lh",
	},
	{
		Name:     "l",
		Expanded: "ls",
	},
}

func TestMatch_NoMatches(t *testing.T) {
	candidate := Match(mockAliasDefs, "cd ..")
	assert.Nil(t, candidate, "should return nil when no matches found")
}

func TestMatch_SingleToken(t *testing.T) {
	candidate := Match(mockAliasDefs, "docker")
	assert.Equal(t, candidate.Name, "dk")
}

func TestMatch_MultipleTokens(t *testing.T) {
	candidate := Match(mockAliasDefs, "git branch")
	assert.Equal(t, candidate.Name, "gb")
}

func TestMatch_MultipleMatches(t *testing.T) {
	candidate := Match(mockAliasDefs, "git checkout -b")
	assert.Equal(t, candidate.Name, "gcb",
		"should return the alias definition that has the longest abbreviation when multiple matches found")
}

func TestMatch_RecursiveDefs(t *testing.T) {
	candidate := Match(mockAliasDefs, "ls -G -lh")
	assert.Equal(t, candidate.Name, "ll", "should apply aliases recursively")
}
