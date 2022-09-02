//go:build prod

package config

func init() {
	GitHubOAuthClientID = "5edf4ed2d423ffa9a133"
	GitHubOAuthCLIToAPIURL = "https://api.yo-lo.sh/github/oauth/callback"
}
