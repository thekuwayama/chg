package lib

import (
	"errors"
	"net/url"
	"testing"
)

func stubHttpGet(us string) (string, error) {
	_, err := url.Parse(us)
	if err != nil {
		return "", errors.New("Fail to parse url")
	}

	return "test", nil
}

func TestWorker(t *testing.T) {
	wp := make(chan chan job, 1)
	w := newWorker(stubHttpGet, wp)
	go w.run()

	jc := <-wp
	d := make(chan string, 1)
	jc <- *newJob("https://example.com", d)

	if "test" != <-d {
		t.Fatal("Not issue stubHttpGet()")
	}
}
