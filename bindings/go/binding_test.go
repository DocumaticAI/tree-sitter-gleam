package tree_sitter_gleam_test

import (
	"testing"

	tree_sitter_gleam "github.com/DocumaticAI/tree-sitter-gleam/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gleam.Language())
	if language == nil {
		t.Errorf("Error loading Gleam grammar")
	}
}
