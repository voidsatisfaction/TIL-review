package main

import (
	"fmt"
	"os"
	"time"

	configreader "github.com/voidsatisfaction/TIL-review/pkg/configReader"
	"github.com/voidsatisfaction/TIL-review/pkg/github"
	"github.com/voidsatisfaction/TIL-review/pkg/mailManager"
	templateengine "github.com/voidsatisfaction/TIL-review/pkg/templateEngine"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: command is not proper")
		fmt.Println("Command should be: ./everydayReview config_file everyday_review_template_file")
		os.Exit(1)
	}
	configFilePath := os.Args[1]
	everydayReviewHTMLFilePath := os.Args[2]

	configFile := configreader.ReadFromJson(configFilePath)

	fmt.Printf("%+v\n", configFile)

	githubClient := github.NewClient(
		configFile.Github.Owner,
		configFile.Github.Repository,
	)

	// Get last 3days commitlist
	// if time.now() == "2019-01-04T10:04:59"
	// since == "2019-01-01T00:00:00"
	since, until := time.Now().AddDate(0, 0, -30).Truncate(24*time.Hour).Add(-9*time.Hour), time.Now()
	last30DaysCommitList := &github.CommitList{}
	for page, perPage := 1, 100; ; page++ {
		commitList, err := githubClient.GetTILCommitList(&since, &until, page, perPage)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}

		slice := append(*last30DaysCommitList, *commitList...)
		last30DaysCommitList = &slice
		if len(slice)%perPage != 0 || len(slice) == 0 {
			break
		}
	}

	last3DaysCommitList := github.CommitList{}
	commitList7DaysAgo := github.CommitList{}
	commitList14DaysAgo := github.CommitList{}
	commitList30DaysAgo := github.CommitList{}
	now := time.Now()
	for _, commit := range *last30DaysCommitList {
		commitDate, err := time.Parse("2006-01-02T15:04:05.9999Z", commit.Commit.Committer.Date)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}

		hoursAgo := int(now.Sub(commitDate).Truncate(24 * time.Hour).Hours())
		if hoursAgo <= 24*3 {
			last3DaysCommitList = append(last3DaysCommitList, commit)
		} else if hoursAgo == 24*7 {
			commitList7DaysAgo = append(commitList7DaysAgo, commit)
		} else if hoursAgo == 24*14 {
			commitList14DaysAgo = append(commitList14DaysAgo, commit)
		} else if hoursAgo == 24*30 {
			commitList30DaysAgo = append(commitList30DaysAgo, commit)
		}
	}

	// Get TIL Master Branch info and TIL Tree
	// branchInfo, err := githubClient.GetBranchInfo("master")
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }

	// tilMasterTreeSha := branchInfo.Commit.Sha

	// FIXME: use tree
	// tree, err := githubClient.GetTILTree(tilMasterTreeSha, true)

	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }

	baseTemplateFilePath := "./template/_base.html.tmpl"
	headerTemplateFilePath := "./template/__header.html.tmpl"
	htmlFilePathList := []string{
		everydayReviewHTMLFilePath,
		baseTemplateFilePath,
		headerTemplateFilePath,
	}

	todayString := time.Now().Format("2006-01-02")

	templateEngine := templateengine.New()
	htmlString, err := templateEngine.CreateHTML(
		htmlFilePathList,
		struct {
			Today               string
			Last3DaysCommitList interface{}
			CommitList7DaysAgo  interface{}
			CommitList14DaysAgo interface{}
			CommitList30DaysAgo interface{}
		}{
			Today:               todayString,
			Last3DaysCommitList: last3DaysCommitList,
			CommitList7DaysAgo:  commitList7DaysAgo,
			CommitList14DaysAgo: commitList14DaysAgo,
			CommitList30DaysAgo: commitList30DaysAgo,
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

	err = mailManager.SendHTML(
		configFile.Mail.From,
		configFile.Mail.To,
		"[TIL-Review] Happy TIL Review Time !!",
		htmlString,
	)

	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
