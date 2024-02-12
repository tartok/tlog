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
		if c == nil {
			continue
		}
		if i == 0 && len(o.c) > 0 {
			c.Write(o.c)
			c.Write(p)
			c.Write([]byte(color.White))
		} else {
			c.Write(p)
		}
	}
	return len(p), nil
}

func GetOut(c string, outs ...io.Writer) io.Writer {
	o := out{c: []byte(c), outs: outs}
	return o
}
