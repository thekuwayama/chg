package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ConncurrentHttpGet(nWorker int) {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("Fail to read stdin")
	}
	urls := strings.Fields(string(in))

	jq := make(chan job, len(urls))
	d := *newDispatcher(nWorker, jq)
	d.run()

	done := make(chan string, len(urls))
	for _, u := range urls {
		jq <- *newJob(u, done)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-done)
	}
}
