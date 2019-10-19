package lib

import (
	"testing"
)

func TestJob(t *testing.T) {
	us := "https://example.com"
	d := make(chan string, 1)
	job := newJob(us, d)

	if job.urlString != us {
		t.Errorf("Got %v expected %v", job.urlString, us)
	}
	if job.done != d {
		t.Errorf("Got %v expected %v", job.done, d)
	}
}
