package main

import "testing"

func TestIsBalanced(t *testing.T) {
	if !IsBalanced(nil) {
		t.Error("empty tree is balanced")
	}
	if !IsBalanced(tn(1, nil, nil)) {
		t.Error("single node is balanced")
	}
	ok := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	if !IsBalanced(ok) {
		t.Error("expected balanced")
	}
	skewed := tn(1, tn(2, tn(3, nil, nil), nil), nil)
	if IsBalanced(skewed) {
		t.Error("left chain of 3 is not balanced")
	}
	// Balanced at the root but not deeper down.
	sneaky := tn(1,
		tn(2, tn(3, tn(4, nil, nil), nil), nil),
		tn(2, nil, tn(3, nil, tn(4, nil, nil))))
	if IsBalanced(sneaky) {
		t.Error("imbalance inside subtrees must be caught")
	}
}
