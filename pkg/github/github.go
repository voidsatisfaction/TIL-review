package github

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
