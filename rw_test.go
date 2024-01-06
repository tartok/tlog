package tlog

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRw(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7}
	r := bytes.NewReader(testData)
	rw := NewReadWriter(func(p []byte) (n int, err error) {
		return r.Read(p)
	}, func(p []byte) (n int, err error) {
		fmt.Println(p)
		return len(p), nil
	})
	p := make([]byte, 2)
	for {
		n, err := rw.Read(p)
		fmt.Println(p[:n], err)
		if err != nil {
			return
		}
		rw.Write(p[:n])
	}
}
