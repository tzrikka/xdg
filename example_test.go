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
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		fmt.Println(err)
		return
	}

	cfgFile := filepath.Join(appDir, "config_file")
	if err := os.WriteFile(cfgFile, []byte("\n"), 0o600); err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleDataHome() {
	rootDir, err := xdg.DataHome()
	if err != nil {
		fmt.Println(err)
		return
	}

	appDir := filepath.Join(rootDir, "my_app")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		fmt.Println(err)
		return
	}

	dataFile := filepath.Join(appDir, "data_file")
	if err := os.WriteFile(dataFile, []byte("\n"), 0o600); err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleMustConfigHome() {
	appDir := filepath.Join(xdg.MustConfigHome(), "my_app")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		fmt.Println(err)
		return
	}

	cfgFile := filepath.Join(appDir, "config_file")
	if err := os.WriteFile(cfgFile, []byte("\n"), 0o600); err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleMustDataHome() {
	appDir := filepath.Join(xdg.MustDataHome(), "my_app")
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		fmt.Println(err)
		return
	}

	dataFile := filepath.Join(appDir, "data_file")
	if err := os.WriteFile(dataFile, []byte("\n"), 0o600); err != nil {
		fmt.Println(err)
		return
	}
}
