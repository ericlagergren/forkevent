package forkevent

import "testing"

func TestPoll(t *testing.T) {
	const (
		path = "/proc/sys/kernel/hostname"
	)
	ch := make(chan struct{})
	err := poll(path, func() bool {
		close(ch)
		return true
	})
	if err != nil {
		t.Fatal(err)
	}
	<-ch
}
