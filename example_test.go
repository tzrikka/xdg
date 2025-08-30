package xdg_test

import (
	"fmt"

	"github.com/tzrikka/xdg"
)

func ExampleConfigHome() {
	path, err := xdg.ConfigHome()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Config home dir for all apps: %s\n", path)
}

func ExampleDataHome() {
	path, err := xdg.DataHome()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Data home dir for all apps: %s\n", path)
}

func ExampleMustConfigHome() {
	fmt.Printf("Config home dir for all apps: %s\n", xdg.MustConfigHome())
}

func ExampleMustDataHome() {
	fmt.Printf("Data home dir for all apps: %s\n", xdg.MustDataHome())
}

func ExampleCreateDir() {
	cfgPath, err := xdg.CreateDir(xdg.ConfigHome, "my_app")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	dataPath, err := xdg.CreateDir(xdg.DataHome, "my_app")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	fmt.Printf("My app's config dir: %s\n", cfgPath)
	fmt.Printf("My app's data dir: %s\n", dataPath)
}

func ExampleCreateFile() {
	cfgPath, err := xdg.CreateFile(xdg.ConfigHome, "my_app", "config_file")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	dataPath, err := xdg.CreateFile(xdg.DataHome, "my_app", "data_file")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	fmt.Printf("My app's existing/new config file path: %s\n", cfgPath)
	fmt.Printf("My app's existing/new data file path: %s\n", dataPath)
}

func ExampleFindConfigFile() {
	path, err := xdg.FindConfigFile("my_app", "config_file")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	if path == "" {
		fmt.Println("File not found")
		return
	}

	fmt.Printf("Found my app's config file in this path: %s\n", path)
}

func ExampleFindDataFile() {
	path, err := xdg.FindDataFile("my_app", "data_file")
	if err != nil {
		fmt.Println(err) // Input or runtime error.
		return
	}

	if path == "" {
		fmt.Println("File not found")
		return
	}

	fmt.Printf("Found my app's data file in this path: %s\n", path)
}
