package main

import (
	"context"
	"github.com/cucumber/godog"
	customassert "github.com/priyankshah217/godog-poc/assertions"
	"strconv"
)

type Result struct{}

func iPerformBelowOnBelowNumbers(ctx context.Context, operation string, table *godog.Table) context.Context {
	op1, _ := strconv.Atoi(table.Rows[1].Cells[0].Value)
	op2, _ := strconv.Atoi(table.Rows[1].Cells[1].Value)
	var actualResult int
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
	return context.WithValue(ctx, Result{}, actualResult)
}

func theShouldBeBelow(ctx context.Context, expectedResult int) error {
	actualResult := ctx.Value(Result{}).(int)
	return customassert.AssertEqual(expectedResult, actualResult)
}
