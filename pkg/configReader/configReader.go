package configreader

import (
	"encoding/json"
	"io/ioutil"
)

/*
example config file format

{
  "github": {
    "username": "voidsatisfaction",
    "repository": "TIL"
  },
  "mail": {
    "from": "cyanicjin1@gmail.com",
    "to": [
      "cyanicjin1@gmail.com"
    ],
    "username": "voidsatisfaction",
    "password": "pass"
  },
  "alarmTime": ["22:00", "23:00"]
}
*/
type configJSON struct {
	Github struct {
		Username   string `json:"username"`
		Repository string `json:"repository"`
	} `json:"github"`
	Mail struct {
		From     string   `json:"from"`
		To       []string `json:"to"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"mail"`
	AlarmTime []string `json:"alarmTime"`
}

func ReadFromJson(filePath string) *configJSON {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	configFile := &configJSON{}
	json.Unmarshal(file, configFile)

	return configFile
}
