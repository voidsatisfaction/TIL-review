package main

import (
	"fmt"
	"os"
	"path"

	configreader "github.com/voidsatisfaction/TIL-review/pkg/configReader"
	"github.com/voidsatisfaction/TIL-review/pkg/mailManager"
	"github.com/voidsatisfaction/TIL-review/pkg/templateEngine"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("executable <path_of_config>")
		fmt.Println("Error: config file path is not set")
		os.Exit(1)
	}
	configFilePath := os.Args[1]
	shortTermHTMLPath := os.Args[2]
	shortTermHTMLFileName := path.Base(shortTermHTMLPath)

	configFile := configreader.ReadFromJson(configFilePath)

	templateEngine := templateEngine.New()
	htmlString, err := templateEngine.CreateHTML(
		shortTermHTMLFileName,
		shortTermHTMLPath,
		struct {
			message string
			test    string
		}{
			message: "hello world!",
			test:    "test argument",
		},
	)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	mailManager := mailManager.New(
		configFile.Mail.Username,
		configFile.Mail.Password,
		configFile.Mail.SmtpHost,
		configFile.Mail.SmtpPort,
	)

	err = mailManager.Send(
		configFile.Mail.From,
		configFile.Mail.To,
		htmlString,
	)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
