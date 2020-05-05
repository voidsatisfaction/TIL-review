package github

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const apiHost string = "https://api.github.com"

type GithubClient struct {
	owner      string
	repository string

	httpClient *resty.Client
}

func NewClient(owner string, repository string) *GithubClient {
	httpClient := resty.New()

	return &GithubClient{
		owner:      owner,
		repository: repository,

		httpClient: httpClient,
	}
}

func newCommitListURLString(owner, repository string) string {
	return fmt.Sprintf("%s/repos/%s/%s/commits", apiHost, owner, repository)
}

type CommitList []Commit

type Commit struct {
	SHA         string `json:"sha,omitempty"`
	HTMLURL     string `json:"html_url,omitempty"`
	URL         string `json:"url,omitempty"`
	CommentsURL string `json:"comments_url,omitempty"`
	Commit      struct {
		Message   string `json:message`
		Committer struct {
			Name  string `json:"name"`
			Date  string `json:"date"`
			Email string `json:"email"`
		} `json:"committer"`
	} `json:commit`
}

func (githubClient *GithubClient) GetTILCommitList(since, until *time.Time, page, perPage int) (*CommitList, error) {
	commitListURLString := newCommitListURLString(
		githubClient.owner,
		githubClient.repository,
	)

	layout := "2006-01-02T15:04:05Z"
	sinceString := since.Format(layout)
	untilString := until.Format(layout)

	resp, err := githubClient.httpClient.R().
		SetQueryParam("since", sinceString).
		SetQueryParam("until", untilString).
		SetQueryParam("page", strconv.Itoa(page)).
		SetQueryParam("per_page", strconv.Itoa(perPage)).
		Get(commitListURLString)

	if err != nil {
		return nil, err
	}

	commitList := &CommitList{}

	err = json.Unmarshal(resp.Body(), commitList)

	if err != nil {
		return nil, err
	}

	return commitList, nil
}

// https://api.github.com/repos/voidsatisfaction/TIL/git/trees/0930bd59a37e93c7243cb63945fe99f8b9eec038?recursive=1
func newGithubTreeURLString(owner, repository, treeSha string) string {
	return fmt.Sprintf("%s/repos/%s/%s/git/trees/%s", apiHost, owner, repository, treeSha)
}

type GithubTree struct {
	Sha  string `json:"sha"`
	Url  string `json:"url"`
	Tree []struct {
		Path string `json:"path"`
	}
}

func (ghc *GithubClient) GetTILTree(treeSha string, recursive bool) (*GithubTree, error) {
	githubTreeURLString := newGithubTreeURLString(ghc.owner, ghc.repository, treeSha)

	recursiveQueryParameter := "0"
	if recursive {
		recursiveQueryParameter = "1"
	}

	resp, err := ghc.httpClient.R().
		SetQueryParam("recursive", recursiveQueryParameter).
		Get(githubTreeURLString)

	if err != nil {
		return nil, err
	}

	githubTree := &GithubTree{}

	err = json.Unmarshal(resp.Body(), githubTree)

	if err != nil {
		return nil, err
	}

	return githubTree, nil
}

func newBranchInfoURLString(owner, repository, branch string) string {
	return fmt.Sprintf("%s/repos/%s/%s/branches/%s", apiHost, owner, repository, branch)
}

type BranchInfo struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
	}
}

// https://api.github.com/repos/voidsatisfaction/TIL/branches/master
func (ghc *GithubClient) GetBranchInfo(branch string) (*BranchInfo, error) {
	branchInfoURLString := newBranchInfoURLString(ghc.owner, ghc.repository, branch)

	resp, err := ghc.httpClient.R().
		Get(branchInfoURLString)

	if err != nil {
		return nil, err
	}

	branchInfo := &BranchInfo{}

	err = json.Unmarshal(resp.Body(), branchInfo)

	if err != nil {
		return nil, err
	}

	return branchInfo, nil
}
