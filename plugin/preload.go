package plugin

import (
	"github.com/lovercode/gopher-lua-libs/argparse"
	"github.com/lovercode/gopher-lua-libs/aws/cloudwatch"
	"github.com/lovercode/gopher-lua-libs/base64"
	"github.com/lovercode/gopher-lua-libs/cert_util"
	"github.com/lovercode/gopher-lua-libs/chef"
	"github.com/lovercode/gopher-lua-libs/cmd"
	"github.com/lovercode/gopher-lua-libs/crypto"
	"github.com/lovercode/gopher-lua-libs/db"
	"github.com/lovercode/gopher-lua-libs/filepath"
	"github.com/lovercode/gopher-lua-libs/goos"
	"github.com/lovercode/gopher-lua-libs/http"
	"github.com/lovercode/gopher-lua-libs/humanize"
	"github.com/lovercode/gopher-lua-libs/inspect"
	"github.com/lovercode/gopher-lua-libs/ioutil"
	"github.com/lovercode/gopher-lua-libs/json"
	"github.com/lovercode/gopher-lua-libs/log"
	"github.com/lovercode/gopher-lua-libs/pb"
	"github.com/lovercode/gopher-lua-libs/pprof"
	prometheus "github.com/lovercode/gopher-lua-libs/prometheus/client"
	"github.com/lovercode/gopher-lua-libs/regexp"
	"github.com/lovercode/gopher-lua-libs/runtime"
	"github.com/lovercode/gopher-lua-libs/shellescape"
	"github.com/lovercode/gopher-lua-libs/stats"
	"github.com/lovercode/gopher-lua-libs/storage"
	"github.com/lovercode/gopher-lua-libs/strings"
	"github.com/lovercode/gopher-lua-libs/tac"
	"github.com/lovercode/gopher-lua-libs/tcp"
	"github.com/lovercode/gopher-lua-libs/telegram"
	"github.com/lovercode/gopher-lua-libs/template"
	"github.com/lovercode/gopher-lua-libs/time"
	"github.com/lovercode/gopher-lua-libs/xmlpath"
	"github.com/lovercode/gopher-lua-libs/yaml"
	"github.com/lovercode/gopher-lua-libs/zabbix"
	lua "github.com/yuin/gopher-lua"
)

// PreloadAll preload all gopher lua packages - note it's needed here to prevent circular deps between plugin and libs
func PreloadAll(L *lua.LState) {
	Preload(L)

	argparse.Preload(L)
	base64.Preload(L)
	cert_util.Preload(L)
	chef.Preload(L)
	cloudwatch.Preload(L)
	cmd.Preload(L)
	crypto.Preload(L)
	db.Preload(L)
	filepath.Preload(L)
	goos.Preload(L)
	http.Preload(L)
	humanize.Preload(L)
	inspect.Preload(L)
	ioutil.Preload(L)
	json.Preload(L)
	log.Preload(L)
	pb.Preload(L)
	pprof.Preload(L)
	prometheus.Preload(L)
	regexp.Preload(L)
	runtime.Preload(L)
	shellescape.Preload(L)
	stats.Preload(L)
	storage.Preload(L)
	strings.Preload(L)
	tac.Preload(L)
	tcp.Preload(L)
	telegram.Preload(L)
	template.Preload(L)
	time.Preload(L)
	xmlpath.Preload(L)
	yaml.Preload(L)
	zabbix.Preload(L)
}
