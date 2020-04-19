package github

import "testing"

type usernameRepositoryExpected struct {
	username   string
	repository string
}

func TestGithubClient(t *testing.T) {
	usernameRepositories := []usernameRepositoryExpected{
		{
			username:   `hello world`,
			repository: `hello repository`,
		},
		{
			username:   `1234`,
			repository: `5678`,
		},
	}

	for _, v := range usernameRepositories {
		githubClient := CreateClient(v.username, v.repository)

		if githubClient.username != v.username {
			t.Fatalf("expected username to be %v, got: %v", v.username, githubClient.username)
		}

		if githubClient.repository != v.repository {
			t.Fatalf("expected repository to be %v, got: %v", v.repository, githubClient.repository)
		}
	}
}
