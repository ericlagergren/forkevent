// Package forkevent allows polling the /proc/sys/kernel/hostname
// to listen for VM forks.
package forkevent

import (
	"os"
)

// Poll invokes fn each time /proc/sys/kernel/random/fork_event
// is updated until fn returns false.
func Poll(fn func() bool) error {
	return poll("/proc/sys/kernel/random/fork_event", fn)
}

func poll(path string, fn func() bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	c, err := f.SyscallConn()
	if err != nil {
		f.Close()
		return err
	}
	go func() {
		defer f.Close()
		c.Read(func(_ uintptr) bool {
			return !fn()
		})
	}()
	return nil
}
