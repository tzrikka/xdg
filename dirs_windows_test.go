package xdg

import (
	"testing"
)

func TestExpandWindows(t *testing.T) {
	t.Setenv("foo", "123")
	t.Setenv("BAR", "456")

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "empty",
			path: "",
			want: "",
		},
		{
			name: "no_replacements",
			path: "abc\\def",
			want: "abc\\def",
		},
		{
			name: "multiple_replacements",
			path: "%foo%;%BAR%",
			want: "123;456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.path); got != tt.want {
				t.Errorf("expand(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
