package zabbix_test

import (
	"github.com/lovercode/gopher-lua-libs/tests"
	"github.com/stretchr/testify/assert"
	"testing"

	http "github.com/lovercode/gopher-lua-libs/http"
	inspect "github.com/lovercode/gopher-lua-libs/inspect"
	zabbix "github.com/lovercode/gopher-lua-libs/zabbix"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		zabbix.Preload,
		http.Preload,
		inspect.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
