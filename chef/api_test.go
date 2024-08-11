package chef_test

import (
	"github.com/lovercode/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	chef "github.com/lovercode/gopher-lua-libs/chef"
	http "github.com/lovercode/gopher-lua-libs/http"
	inspect "github.com/lovercode/gopher-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		chef.Preload,
		http.Preload,
		inspect.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
