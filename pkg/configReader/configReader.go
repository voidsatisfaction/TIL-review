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
    "from": "youremail@gmail.com",
    "to": [
	  "youremail@gmail.com"
    ],
    "username": "voidsatisfaction",
    "password": "pass"
  },
  "alarmTime": ["22:00", "23:00"]
}
*/
type configJSON struct {
	Github struct {
		Owner      string `json:"owner"`
		Repository string `json:"repository"`
		Branch     string `json:"branch"`
	} `json:"github"`
	Mail struct {
		SmtpHost string   `json:"smtpHost"`
		SmtpPort string   `json:"smtpPort"`
		From     string   `json:"from"`
		To       []string `json:"to"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"mail"`
	AlarmTime []string `json:"alarmTime"`
	Template  struct {
		EveryDayReview struct {
			RandomPicks             int      `json:"randomPicks"`
			PinnedFileOrFolderPaths []string `json:"pinnedFileOrFolderPaths"`
			LastNDaysCommits        int      `json:"lastNDaysCommits"`
			CommitListNDaysAgo      []int    `json:"commitListNDaysAgo"`
		} `json:"everydayReview"`
	} `json:"template"`
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
