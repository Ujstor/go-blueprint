package template

import (
	_ "embed"
)

//go:embed files/routes/fiber.go.tmpl
var fiberRoutesTemplate []byte

//go:embed files/server/fiber.go.tmpl
var fiberServerTemplate []byte

//go:embed files/main/fiber_main.go.tmpl
var fiberMainTemplate []byte

//go:embed files/tests/server/server_test_fiber.go.tmpl
var fiberServerTestTemplate []byte

//go:embed files/tests/routes/routes_test_fiber.go.tmpl
var fiberRoutesTestTemplate []byte

// FiberTemplates contains the methods used for building
// an app that uses [github.com/gofiber/fiber]
type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return fiberMainTemplate
}
func (f FiberTemplates) Server() []byte {
	return fiberServerTemplate
}

func (f FiberTemplates) Routes() []byte {
	return fiberRoutesTemplate
}

func (f FiberTemplates) ServerTest() []byte {
	return fiberServerTestTemplate
}

func (f FiberTemplates) RoutesTest() []byte {
	return fiberRoutesTestTemplate
}