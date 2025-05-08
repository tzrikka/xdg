package xdg

import (
	"os"
	"testing"
)

func TestExpandUnix(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "empty",
			path: "",
			want: ".",
		},
		{
			name: "minimum",
			path: "~",
			want: "${HOME}",
		},
		{
			name: "with_suffix",
			path: "~/foo",
			want: "${HOME}/foo",
		},
		{
			name: "only_first_char",
			path: "foo/~/bar",
			want: "foo/~/bar",
		},
		{
			name: "only_once",
			path: "~/foo/~/bar",
			want: "${HOME}/foo/~/bar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := os.ExpandEnv(tt.want)
			if got := expand(tt.path); got != want {
				t.Errorf("expand(%q) = %q, want %q", tt.path, got, want)
			}
		})
	}
}
