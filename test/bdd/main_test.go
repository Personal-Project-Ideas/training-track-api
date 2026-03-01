// Package bdd contains BDD test scenarios for the application.
package bdd

import (
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Format:   "pretty",
		Paths:    []string{"../features"},
		TestingT: t,
	}
	godog.TestSuite{
		Name:                "user",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()
}
