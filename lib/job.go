package lib

type job struct {
	urlString string
	done      chan string
}

func newJob(us string, d chan string) *job {
	return &job{urlString: us, done: d}
}
