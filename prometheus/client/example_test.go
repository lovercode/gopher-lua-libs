package prometheus_client_test

import (
	"log"

	"github.com/lovercode/gopher-lua-libs/http"
	prometheus "github.com/lovercode/gopher-lua-libs/prometheus/client"
	"github.com/lovercode/gopher-lua-libs/time"
	lua "github.com/yuin/gopher-lua"
)

// prometheus:start(string)
func ExampleStart() {
	state := lua.NewState()
	prometheus.Preload(state)
	time.Preload(state)
	http.Preload(state)

	source := `
    local prometheus = require("prometheus")
	local time = require("time")
	local http = require("http_client")

	local pp = prometheus.register(":18080")
	pp:start()
    time.sleep(1)

	local client = http.client({timeout=5})

	local request = http.request("GET", "http://127.0.0.1:18080/")
	local result = client:do_request(request)
	print(result.code)

	local request = http.request("GET", "http://127.0.0.1:18080/metrics")
	local result = client:do_request(request)
	print(result.code)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// 404
	// 200
}
