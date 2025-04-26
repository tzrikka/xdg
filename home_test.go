package xdg

import (
	"testing"
)

func TestHomeDir(t *testing.T) {
	t.Cleanup(func() { cachedHomeDir = "" })
	want := "TEST1"

	tests := []struct {
		name string
		env  string
	}{
		{
			name: "without_cache",
			env:  want,
		},
		{
			name: "with_cache",
			env:  "TEST2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("HOME", tt.env)
			t.Setenv("USERPROFILE", tt.env)
			if got := HomeDir(); got != want {
				t.Errorf("HomeDir() = %q, want %q", got, want)
			}
		})
	}
}
