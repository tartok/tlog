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
	FuncWriter struct {
		w []func(p []byte) (n int, err error)
	}
	FuncReader struct {
		r []func(p []byte) (n int, err error)
	}
)

func NewFuncWriter(f ...func(p []byte) (n int, err error)) io.Writer {
	return FuncWriter{w: f}
}
func (f FuncWriter) Write(p []byte) (n int, err error) {
	for _, f := range f.w {
		f(p)
	}
	return len(p), nil
}
func NewFuncReader(f ...func(p []byte) (n int, err error)) io.Reader {
	return FuncReader{r: f}
}
func (f FuncReader) Read(p []byte) (n int, err error) {
	for _, f := range f.r {
		f(p)
	}
	return len(p), nil
}
func (o out) Write(p []byte) (n int, err error) {
	if len(o.outs) == 0 {
		return len(p), nil
	}
	for i, c := range o.outs {
		if c == nil {
			continue
		}
		if i == 0 && len(o.c) > 0 {
			c(o.c)
			n, err = c(p)
			c([]byte(color.White))
		} else {
			n, err = c(p)
		}
	}
	return
}

func GetOut(c string, outs ...func(p []byte) (n int, err error)) io.Writer {
	o := out{c: []byte(c), outs: outs}
	return o
}
