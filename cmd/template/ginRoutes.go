package template

import (
	_ "embed"
)

//go:embed files/routes/gin.go.tmpl
var ginRoutesTemplate []byte

//go:embed files/tests/routes/routes_test_gin.go.tmpl
var ginRoutesTestTemplate []byte

// GinTemplates contains the methods used for building
// an app that uses [github.com/gin-gonic/gin]
type GinTemplates struct{}

func (g GinTemplates) Main() []byte {
	return mainTemplate
}

func (g GinTemplates) Server() []byte {
	return standardServerTemplate
}

func (g GinTemplates) Routes() []byte {
	return ginRoutesTemplate
}

func (g GinTemplates) ServerTest() []byte {
	return standardServerTestTemplate
}

func (g GinTemplates) RoutesTest() []byte {
	return ginRoutesTestTemplate
}