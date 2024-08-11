package pprof_test

import (
	lua_http "github.com/lovercode/gopher-lua-libs/http"
	lua_pprof "github.com/lovercode/gopher-lua-libs/pprof"
	"github.com/lovercode/gopher-lua-libs/tests"
	lua_time "github.com/lovercode/gopher-lua-libs/time"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		lua_pprof.Preload,
		lua_http.Preload,
		lua_time.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
