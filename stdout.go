package tlog

import (
	"github.com/TwiN/go-color"
	"io"
)

type (
	out struct {
		c      []byte
		outs   []io.Writer
		prefix []byte
	}
)

func (o out) Write(p []byte) (n int, err error) {
	if len(o.outs) == 0 {
		return len(p), nil
	}
	for i, c := range o.outs {
		if i == 0 {
			c.Write(o.c)
			n, err = c.Write(p)
			c.Write([]byte(color.White))
		} else {
			c.Write(p)
		}
	}
	return
}

func getOut(c string, outs ...io.Writer) io.Writer {
	o := out{c: []byte(c), outs: outs}
	return o
}
func drrorOut() io.Writer {
	o := out{c: []byte(color.Red)}
	return o
}
