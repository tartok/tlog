package tlog

import (
	"github.com/TwiN/go-color"
	"io"
)

type (
	out struct {
		c      []byte
		outs   []func(p []byte) (n int, err error)
		prefix []byte
	}
)

func (o out) Write(p []byte) (n int, err error) {
	if len(o.outs) == 0 {
		return len(p), nil
	}
	for i, c := range o.outs {
		if i == 0 {
			c(o.c)
			n, err = c(p)
			c([]byte(color.White))
		} else {
			c(p)
		}
	}
	return
}

func GetOut(c string, outs ...func(p []byte) (n int, err error)) io.Writer {
	o := out{c: []byte(c), outs: outs}
	return o
}
func drrorOut() io.Writer {
	o := out{c: []byte(color.Red)}
	return o
}
