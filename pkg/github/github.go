package github

import (
	"github.com/go-resty/resty/v2"
)

type githubClient struct {
	username   string
	repository string

	httpClient *resty.Client
}

func CreateClient(username string, repository string) *githubClient {
	httpClient := resty.New()

	return &githubClient{
		username:   username,
		repository: repository,

		httpClient: httpClient,
	}
}
