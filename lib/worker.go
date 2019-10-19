package lib

type worker struct {
	clientFunc func(string) (string, error)
	jobChannel chan job
	workerPool chan chan job
}

func newWorker(f func(string) (string, error), wp chan chan job) *worker {
	return &worker{clientFunc: f, workerPool: wp, jobChannel: make(chan job, 1)}
}

func (w *worker) run() {
	for {
		w.workerPool <- w.jobChannel
		j := <-w.jobChannel

		httpBody, err := w.clientFunc(j.urlString)
		if err != nil {
			j.done <- err.Error()
		} else {
			j.done <- httpBody
		}
	}
}
