package github

import (
	"slices"
	"testing"
)

func TestClassifyLanguages(t *testing.T) {
	langs := map[string]int{
		"Go":         45000,
		"Dockerfile": 200,
		"HCL":        1500,
		"Rego":       300, // linguagem não mapeada, deveria cair em Other
	}

	stack := ClassifyLanguages(langs)

	if !slices.Contains(stack.Languages, "Go") {
		t.Errorf("expected Languages to contain Go, got %v", stack.Languages)
	}

	if !slices.Contains(stack.Infra, "Dockerfile") {
		t.Errorf("expected Infra to contain Dockerfile, got %v", stack.Infra)
	}

	if !slices.Contains(stack.Infra, "HCL") {
		t.Errorf("expected Infra to contain HCL, got %v", stack.Infra)
	}

	if !slices.Contains(stack.Other, "Rego") {
		t.Errorf("expected Other to contain Rego, got %v", stack.Other)
	}

	if len(stack.Languages) != 1 {
		t.Errorf("expected 1 language, got %d", len(stack.Languages))
	}

	if len(stack.Infra) != 2 {
		t.Errorf("expected 2 infra items, got %d", len(stack.Infra))
	}

	if len(stack.Other) != 1 {
		t.Errorf("expected 1 other item, got %d", len(stack.Other))
	}
}
