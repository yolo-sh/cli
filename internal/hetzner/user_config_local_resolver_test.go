package hetzner

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yolo-sh/cli/internal/mocks"
	"github.com/yolo-sh/hetzner-cloud-provider/userconfig"
)

func TestUserConfigLocalResolving(t *testing.T) {
	testCases := []struct {
		test                            string
		configInEnvVars                 *userconfig.Config
		errorDuringEnvVarsResolving     error
		configInFiles                   *userconfig.Config
		errorDuringConfigFilesResolving error
		contextOpts                     string
		expectedConfig                  *userconfig.Config
		expectedError                   error
	}{
		{
			test:                            "no env vars, no config files",
			errorDuringEnvVarsResolving:     userconfig.ErrMissingConfig,
			errorDuringConfigFilesResolving: userconfig.ErrMissingConfig,
			expectedConfig:                  nil,
			expectedError:                   userconfig.ErrMissingConfig,
		},

		{
			test:                            "only env vars",
			configInEnvVars:                 userconfig.NewConfig("a", "b"),
			errorDuringConfigFilesResolving: userconfig.ErrMissingConfig,
			expectedConfig:                  userconfig.NewConfig("a", "b"),
			expectedError:                   nil,
		},

		{
			test:                        "only config files",
			errorDuringEnvVarsResolving: userconfig.ErrMissingConfig,
			configInFiles:               userconfig.NewConfig("a", "b"),
			expectedConfig:              userconfig.NewConfig("a", "b"),
			expectedError:               nil,
		},

		{
			test:            "env vars and config files",
			configInEnvVars: userconfig.NewConfig("a", "b"),
			configInFiles:   userconfig.NewConfig("c", "d"),
			expectedConfig:  userconfig.NewConfig("a", "b"),
			expectedError:   nil,
		},

		{
			test:            "env vars, config files and context",
			configInEnvVars: userconfig.NewConfig("a", "b"),
			configInFiles:   userconfig.NewConfig("c", "d"),
			contextOpts:     "production",
			expectedConfig:  userconfig.NewConfig("c", "d"),
			expectedError:   nil,
		},

		{
			test:                            "env vars and errored config files",
			configInEnvVars:                 userconfig.NewConfig("a", "b"),
			errorDuringConfigFilesResolving: userconfig.ErrMissingRegion,
			expectedConfig:                  userconfig.NewConfig("a", "b"),
			expectedError:                   nil,
		},

		{
			test:                            "env vars, errored config files and context",
			configInEnvVars:                 userconfig.NewConfig("a", "b"),
			errorDuringConfigFilesResolving: userconfig.ErrMissingRegion,
			contextOpts:                     "production",
			expectedConfig:                  nil,
			expectedError:                   userconfig.ErrMissingRegion,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userConfigEnvVarsResolverMock := mocks.NewHetznerUserConfigEnvVarsResolver(mockCtrl)
			userConfigEnvVarsResolverMock.EXPECT().Resolve().Return(
				tc.configInEnvVars,
				tc.errorDuringEnvVarsResolving,
			).AnyTimes()

			userConfigFilesResolverMock := mocks.NewHetznerUserConfigFilesResolver(mockCtrl)
			userConfigFilesResolverMock.EXPECT().Resolve().Return(
				tc.configInFiles,
				tc.errorDuringConfigFilesResolving,
			).AnyTimes()

			resolver := NewUserConfigLocalResolver(
				userConfigEnvVarsResolverMock,
				userConfigFilesResolverMock,
				UserConfigLocalResolverOpts{
					Context: tc.contextOpts,
				},
			)

			resolvedConfig, err := resolver.Resolve()

			if tc.expectedError == nil && err != nil {
				t.Fatalf("expected no error, got '%+v'", err)
			}

			if tc.expectedError != nil && !errors.Is(err, tc.expectedError) {
				t.Fatalf("expected error to equal '%+v', got '%+v'", tc.expectedError, err)
			}

			if tc.expectedConfig != nil && *resolvedConfig != *tc.expectedConfig {
				t.Fatalf("expected config to equal '%+v', got '%+v'", *tc.expectedConfig, *resolvedConfig)
			}

			if tc.expectedConfig == nil && resolvedConfig != nil {
				t.Fatalf("expected no config, got '%+v'", *resolvedConfig)
			}
		})
	}
}
