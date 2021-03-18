package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// GIVEN
	want := "世界マン"
	// WHEN
	if got := tesuto();
	// THEN
	got != want {
		t.Errorf("tesuto() = %q, want %q", got, want)
	}
}
