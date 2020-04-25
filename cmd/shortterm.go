package main

import (
	"fmt"
	"os"

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
	shorttermHTMLPath := os.Args[2]

	configFile := configreader.ReadFromJson(configFilePath)

	fmt.Printf("%+v", configFile)

	templateEngine := templateEngine.New()
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	htmlString, err := templateEngine.CreateHTML(
		shorttermHTMLPath,
		struct{ message string }{message: "hello world!"},
	)
	htmlString = mime + htmlString

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
