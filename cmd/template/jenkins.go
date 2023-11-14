package template

import (
	_ "embed"
)

//go:embed files/cicd/jenkins_pipline.tmpl
var jenkinsTemplate []byte

//go:embed files/cicd/docker_tag.sh.tmpl
var dockerTagTemplate []byte

//go:embed files/cicd/dockerfile_go.tmpl
var dockerfileTemplate []byte

type JenkinsTemplate struct{}

func (j JenkinsTemplate) Pipline() []byte {
	return gitHubActionBuildTemplate
}

func (j JenkinsTemplate) Pipline2() []byte {
	return gitHubActionTestTemplate
}

func (j JenkinsTemplate) Pipline3() []byte {
	return gitHubActionLintingTemplate
}