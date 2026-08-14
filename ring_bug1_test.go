package main

import "testing"

// TestGetNodeTrimConsistency verifies that GetNode handles whitespace
// in node names the same way AddNode does (trim before lookup).
func TestGetNodeTrimConsistency(t *testing.T) {
	s := NewService()
	// AddNode trims "  mynode  " to "mynode" internally
	if _, err := s.AddNode("  mynode  ", 10); err != nil {
		t.Fatalf("AddNode: %v", err)
	}

	// GetNode with the same untrimmed input must find the node
	node, err := s.GetNode("  mynode  ")
	if err != nil {
		t.Fatalf("GetNode with spaces should find the node after trimming: %v", err)
	}
	if node.Name != "mynode" {
		t.Fatalf("GetNode returned name=%q, want %q", node.Name, "mynode")
	}
}
