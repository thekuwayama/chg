package cli

import (
	"flag"
	"runtime"

	"github.com/thekuwayama/chg/concurrent_http_get"
)

var nWorker = flag.Int("n", runtime.NumCPU(), "number of workers")

func Run() {
	flag.Parse()

	lib.ConncurrentHttpGet(*nWorker)
}
