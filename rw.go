package tlog

type (
	Writer struct {
		w func(p []byte) (n int, err error)
	}
	Reader struct {
		r func(p []byte) (n int, err error)
	}
	ReadWriter struct {
		w func(p []byte) (n int, err error)
		r func(p []byte) (n int, err error)
	}
)

func (r ReadWriter) Read(p []byte) (n int, err error) {
	return r.r(p)
}

func (r ReadWriter) Write(p []byte) (n int, err error) {
	return r.w(p)
}

func (r Reader) Read(p []byte) (n int, err error) {
	return r.r(p)
}

func (w Writer) Write(p []byte) (n int, err error) {
	return w.w(p)
}
func NewWriter(w func(p []byte) (n int, err error)) *Writer {
	return &Writer{w: w}
}
func NewReader(r func(p []byte) (n int, err error)) *Reader {
	return &Reader{r: r}
}
func NewReadWriter(r, w func(p []byte) (n int, err error)) *ReadWriter {
	return &ReadWriter{r: r, w: w}
}
