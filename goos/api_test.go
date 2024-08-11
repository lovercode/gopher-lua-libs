package goos

import (
	"github.com/lovercode/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	runtime "github.com/lovercode/gopher-lua-libs/runtime"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		runtime.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
