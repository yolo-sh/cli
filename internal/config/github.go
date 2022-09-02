package config

var (
	GitHubOAuthClientID    = "dfa945a930f44d9be359"
	GitHubOAuthCLIToAPIURL = "http://127.0.0.1:8080/github/oauth/callback"

	GitHubOAuthAPIToCLIURLPath = "/github/oauth/callback"

	GitHubOAuthScopes = []string{
		"read:user",
		"user:email",
		"repo",
		"admin:public_key",
		"admin:gpg_key",
	}
)
