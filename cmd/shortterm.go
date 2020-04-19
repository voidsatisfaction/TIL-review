package main

import (
	"fmt"
	"os"

	configreader "github.com/voidsatisfaction/TIL-review/pkg/configReader"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("executable <path_of_config>")
		fmt.Println("Error: config file path is not set")
		os.Exit(1)
	}
	configFilePath := os.Args[1]

	configFile := configreader.ReadFromJson(configFilePath)

	fmt.Printf("%+v", configFile)
}
