package template

import (
	_ "embed"
)

//go:embed files/cicd/github_action_goreliser.yml.tmpl
var gitHubActionBuildTemplate []byte

//go:embed files/cicd/github_action_gotest.yml.tmpl
var gitHubActionTestTemplate []byte

//go:embed files/cicd/github_action_linting.yml.tmpl
var gitHubActionLintingTemplate []byte

type GitHubActionTemplate struct{}

func (a GitHubActionTemplate) Pipline1() []byte {
	return gitHubActionBuildTemplate
}

func (a GitHubActionTemplate) Pipline2() []byte {
	return gitHubActionTestTemplate
}

func (a GitHubActionTemplate) Pipline3() []byte {
	return gitHubActionLintingTemplate
}