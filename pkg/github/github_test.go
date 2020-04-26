package github

import (
	"fmt"
	"testing"
	"time"
)

type ownerRepositoryExpected struct {
	owner      string
	repository string
}

func TestGithubClient(t *testing.T) {
	ownerRepositories := []ownerRepositoryExpected{
		{
			owner:      `hello world`,
			repository: `hello repository`,
		},
		{
			owner:      `1234`,
			repository: `5678`,
		},
	}

	for _, v := range ownerRepositories {
		githubClient := NewClient(v.owner, v.repository)

		if githubClient.owner != v.owner {
			t.Fatalf("expected owner to be %v, got: %v", v.owner, githubClient.owner)
		}

		if githubClient.repository != v.repository {
			t.Fatalf("expected repository to be %v, got: %v", v.repository, githubClient.repository)
		}
	}
}

func TestGithubClientGetTILCommits(t *testing.T) {
	owner, repository := "voidsatisfaction", "TIL"
	ghc := NewClient(owner, repository)
	since, until := time.Now().AddDate(0, 0, -3), time.Now()

	commitList, err := ghc.GetTILCommitList(&since, &until, 1, 5)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	for _, commit := range *commitList {
		fmt.Println(commit)
	}
}
