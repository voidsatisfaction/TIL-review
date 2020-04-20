package github

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const apiHost string = "https://api.github.com"

func newCommitListURLString(owner, repository string) string {
	return fmt.Sprintf("%s/repos/%s/%s/commits", apiHost, owner, repository)
}

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

type CommitList []Commit

type Commit struct {
	SHA         string `json:"sha,omitempty"`
	HTMLURL     string `json:"html_url,omitempty"`
	URL         string `json:"url,omitempty"`
	CommentsURL string `json:"comments_url,omitempty"`
}

func (githubClient *GithubClient) GetCommitList() (*CommitList, error) {
	commitListURLString := newCommitListURLString(
		githubClient.owner,
		githubClient.repository,
	)

	resp, err := githubClient.httpClient.R().
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
