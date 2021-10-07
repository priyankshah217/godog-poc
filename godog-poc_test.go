package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
)

func Test(t *testing.T) {
	godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
	}.Run()
}

var actualResult = 0

func iPerformBelowOnBelowNumbers(operation string, table *godog.Table) error {
	op1, _ := strconv.Atoi(table.Rows[1].Cells[0].Value)
	op2, _ := strconv.Atoi(table.Rows[1].Cells[1].Value)
	switch operation {
	case "Add":
		actualResult = op1 + op2
		break
	case "Min":
		actualResult = op1 - op2
		break
	case "Mul":
		actualResult = op1 * op2
		break
	case "Div":
		actualResult = op1 / op2
		break
	}
	return nil
}

func theShouldBeBelow(expectedResult int) error {
	if actualResult != expectedResult {
		return fmt.Errorf("Assertion failed: Expected %d, Actual %d ", expectedResult, actualResult)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I perform below "([^"]*)" on below numbers$`, iPerformBelowOnBelowNumbers)
	ctx.Step(`^the (\d+) should be below$`, theShouldBeBelow)
}
