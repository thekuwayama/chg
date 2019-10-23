package lib

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type dispatcher struct {
	workerPool chan chan job
	nWorker    int
	jobQueue   chan job
}

func newDispatcher(n int, jq chan job) *dispatcher {
	pool := make(chan chan job, n)
	return &dispatcher{workerPool: pool, nWorker: n, jobQueue: jq}
}

func httpGet(us string) (string, error) {
	u, err := url.Parse(us)
	if err != nil {
		return "", errors.New("Fail to parse url")
	}

	res, err := http.Get(u.String())
	if err != nil {
		return "", errors.New("Fail to do the HTTP GET")
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Fail to read HTTP body")
	}
	return string(b), nil
}

func (d *dispatcher) dispatch() {
	for {
		j := <-d.jobQueue
		jc := <-d.workerPool
		jc <- j
	}
}

func (d *dispatcher) run() {
	for i := 0; i < d.nWorker; i++ {
		w := newWorker(httpGet, d.workerPool)
		go w.run()
	}
	go d.dispatch()
}
