package tlog

type (
	Writer struct {
		w []func(p []byte) (n int, err error)
	}
	Reader struct {
		r []func(p []byte) (n int, err error)
	}
	ReadWriter struct {
		w []func(p []byte) (n int, err error)
		r func(p []byte) (n int, err error)
	}
)

func (r ReadWriter) Read(p []byte) (n int, err error) {
	return r.r(p)
}

func (r ReadWriter) Write(p []byte) (n int, err error) {
	for _, f := range r.w {
		f(p)
	}
	return len(p), nil
}

func (r Reader) Read(p []byte) (n int, err error) {
	for _, f := range r.r {
		c, err := f(p)
		n += c
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (w Writer) Write(p []byte) (n int, err error) {
	for _, f := range w.w {
		f(p)
	}
	return len(p), nil
}
func NewWriter(w ...func(p []byte) (n int, err error)) *Writer {
	return &Writer{w: w}
}
func NewReader(r ...func(p []byte) (n int, err error)) *Reader {
	return &Reader{r: r}
}
func NewReadWriter(r func(p []byte) (n int, err error), w ...func(p []byte) (n int, err error)) *ReadWriter {
	return &ReadWriter{r: r, w: w}
}
