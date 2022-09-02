package interfaces

import (
	"github.com/yolo-sh/cli/internal/config"
	"github.com/yolo-sh/yolo/github"
)

type UserConfigManager interface {
	GetString(key config.UserConfigKey) string
	GetBool(key config.UserConfigKey) bool
	Set(key config.UserConfigKey, value interface{})
	WriteConfig() error
	PopulateFromGitHubUser(githubUser *github.AuthenticatedUser)
}
