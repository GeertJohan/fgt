package main

import (
	"io"
	"os"
	"os/exec"
)

var anything bool

type fgtWriter struct {
	wr io.Writer
}

func (fgt *fgtWriter) Write(b []byte) (int, error) {
	if len(b) > 0 {
		anything = true
	}
	return fgt.wr.Write(b)
}

// fast go tester
func main() {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &fgtWriter{
		wr: os.Stdout,
	}
	cmd.Stderr = &fgtWriter{
		wr: os.Stderr,
	}
	cmd.Run()
	if anything {
		os.Exit(1)
	}
}
