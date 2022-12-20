///go:build unit_test

package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnv(t *testing.T) {
	res := ReadEnv("environment")
	err := ReadEnv("")
	assert.Equal(t, "success", res)
	assert.Equal(t, "failed", err)
}
