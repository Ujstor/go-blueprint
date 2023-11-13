package template

import (
	_ "embed"
)

//go:embed files/routes/chi.go.tmpl
var chiRoutesTemplate []byte

//go:embed files/tests/routes/routes_test_chi.go.tmpl
var chiRoutesTestTemplate []byte

// ChiTemplates contains the methods used for building
// an app that uses [github.com/go-chi/chi]
type ChiTemplates struct{}

func (c ChiTemplates) Main() []byte {
	return mainTemplate
}

func (c ChiTemplates) Server() []byte {
	return standardServerTemplate
}

func (c ChiTemplates) Routes() []byte {
	return chiRoutesTemplate
}

func (c ChiTemplates) ServerTest() []byte {
	return standardServerTestTemplate
}

func (c ChiTemplates) RoutesTest() []byte {
	return chiRoutesTestTemplate
}