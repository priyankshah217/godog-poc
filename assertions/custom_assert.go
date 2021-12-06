package assertions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

func AssertEqual(expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	assert.Equal(&t, expected, actual, msgAndArgs...)
	return t.err
}

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
