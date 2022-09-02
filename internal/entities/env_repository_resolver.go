package entities

import (
	// "github.com/fatih/color"
	// "github.com/google/go-github/v43/github"
	"github.com/yolo-sh/cli/internal/config"
	"github.com/yolo-sh/cli/internal/interfaces"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/github"
)

type EnvRepositoryResolver struct {
	logger     interfaces.Logger
	userConfig interfaces.UserConfigManager
	github     interfaces.GitHubManager
}

func NewEnvRepositoryResolver(
	logger interfaces.Logger,
	userConfig interfaces.UserConfigManager,
	github interfaces.GitHubManager,
) EnvRepositoryResolver {

	return EnvRepositoryResolver{
		logger:     logger,
		userConfig: userConfig,
		github:     github,
	}
}

func (e EnvRepositoryResolver) Resolve(
	repositoryName string,
	checkForRepositoryExistence bool,
) (*entities.ResolvedEnvRepository, error) {

	githubAccessToken := e.userConfig.GetString(
		config.UserConfigKeyGitHubAccessToken,
	)

	githubUsername := e.userConfig.GetString(
		config.UserConfigKeyGitHubUsername,
	)

	parsedRepoName, err := github.ParseRepositoryName(
		repositoryName,
		githubUsername,
	)

	if err != nil {
		// If repository name is invalid, we are sure
		// that the repository doesn't exist.
		return nil, entities.ErrEnvRepositoryNotFound{
			RepoOwner: githubUsername,
			RepoName:  repositoryName,
		}
	}

	if checkForRepositoryExistence {
		repoExists, err := e.github.DoesRepositoryExist(
			githubAccessToken,
			parsedRepoName.Owner,
			parsedRepoName.Name,
		)

		if err != nil {
			return nil, err
		}

		// if !repoExists && parsedRepoName.Owner != githubUsername {
		if !repoExists {
			return nil, entities.ErrEnvRepositoryNotFound{
				RepoOwner: parsedRepoName.Owner,
				RepoName:  parsedRepoName.Name,
			}
		}
	}

	// if !repoExists {
	// 	bold := color.New(color.Bold).SprintFunc()

	// 	d.logger.Log(
	// 		"\n%s "+bold("Repository \"%s\" not found. Creating now..."),
	// 		bold(color.YellowString("Warning!")),
	// 		parsedRepoName.Name,
	// 	)

	// 	// Means that we want the repository to be created
	// 	// in the logged user personal account. See GitHub SDK docs.
	// 	createdRepoOrganization := ""

	// 	createdRepoIsPrivate := true
	// 	createdRepoProps := &github.Repository{
	// 		Name:    &parsedRepoName.Name,
	// 		Private: &createdRepoIsPrivate,
	// 	}

	// 	_, err := d.github.CreateRepository(
	// 		githubAccessToken,
	// 		createdRepoOrganization,
	// 		createdRepoProps,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &entities.ResolvedEnvRepository{
		Owner:         parsedRepoName.Owner,
		ExplicitOwner: parsedRepoName.ExplicitOwner,

		Name: parsedRepoName.Name,

		GitURL: github.BuildGitURL(
			parsedRepoName.Owner,
			parsedRepoName.Name,
		),

		GitHTTPURL: github.BuildGitHTTPURL(
			parsedRepoName.Owner,
			parsedRepoName.Name,
		),
	}, nil
}
