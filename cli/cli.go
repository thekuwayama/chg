package cli

import (
	"flag"
	"runtime"

	"github.com/thekuwayama/chg/lib"
)

var nWorker = flag.Int("n", runtime.NumCPU(), "number of workers")

func Run() {
	flag.Parse()

	lib.ConncurrentHttpGet(*nWorker)
}
