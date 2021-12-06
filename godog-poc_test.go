package main

import (
	"testing"

	"github.com/cucumber/godog"
)

func Test(t *testing.T) {
	godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
	}.Run()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I perform below "([^"]*)" on below numbers$`, iPerformBelowOnBelowNumbers)
	ctx.Step(`^the (\d+) should be below$`, theShouldBeBelow)
}
