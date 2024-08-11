package cloudwatch

import (
	"github.com/lovercode/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/lovercode/gopher-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
