package prometheus_client_test

import (
	"github.com/lovercode/gopher-lua-libs/http"
	prometheus "github.com/lovercode/gopher-lua-libs/prometheus/client"
	"github.com/lovercode/gopher-lua-libs/strings"
	"github.com/lovercode/gopher-lua-libs/tests"
	"github.com/lovercode/gopher-lua-libs/time"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		prometheus.Preload,
		http.Preload,
		strings.Preload,
		time.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
