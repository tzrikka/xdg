package xdg_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tzrikka/xdg"
)

func ExampleConfigHome() {
	rootDir, err := xdg.ConfigHome()
	if err != nil {
		fmt.Println(err)
		return
	}

	appDir := filepath.Join(rootDir, "my_app")
	os.MkdirAll(appDir, 0o700)

	cfgFile := filepath.Join(appDir, "config.yaml")
	os.WriteFile(cfgFile, []byte("#\n"), 0o600)
}

func ExampleMustHome() {
	path1 := xdg.MustHome(xdg.ConfigHome())
	fmt.Println(path1)

	path2, err := xdg.ConfigHome()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path2)
}

func ExampleMustDirs() {
	paths1 := xdg.MustDirs(xdg.DataDirs())
	fmt.Println(paths1)

	paths2, err := xdg.DataDirs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(paths2)
}
