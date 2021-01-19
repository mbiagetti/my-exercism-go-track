package paasio

import (
	"io"
	"sync"
)

// ReadWriteStat implement a ReadWrite with statistics
type ReadWriteStat struct {
	sync.RWMutex
	w        io.Writer
	r        io.Reader
	totRead  int64
	totWrite int64
	readOps  int
	writeOps int
}

// Read reads the next len(p) bytes from the buffer or until the buffer
// is drained. The return value n is the number of bytes read. If the
// buffer has no data to return, err is io.EOF (unless len(p) is zero);
// otherwise it is nil.
func (r *ReadWriteStat) Read(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	r.readOps++
	if n, err = r.r.Read(p); err != nil {
		return 0, err
	}
	r.totRead += int64(n)
	return n, nil
}

// Write appends the contents of p to the buffer, growing the buffer as
// needed. The return value n is the length of p; err is always nil. If the
// buffer becomes too large, Write will panic with ErrTooLarge.
func (r *ReadWriteStat) Write(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	r.writeOps++
	if n, err = r.w.Write(p); err != nil {
		return 0, err
	}
	r.totWrite += int64(n)
	return n, nil
}

// ReadCount returns the total number of bytes successfully read along
// with the total number of calls to Read().
func (r *ReadWriteStat) ReadCount() (n int64, nops int) {
	r.RLock()
	defer r.RUnlock()

	return r.totRead, r.readOps
}

// WriteCount returns the total number of bytes successfully written along
// with the total number of calls to Write().
func (r *ReadWriteStat) WriteCount() (n int64, nops int) {
	r.RLock()
	defer r.RUnlock()

	return r.totWrite, r.writeOps
}

// NewWriteCounter return a WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &ReadWriteStat{w: w}
}

// NewReadCounter return a ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &ReadWriteStat{r: r}
}

// NewReadWriteCounter return a ReadWriteCounter
func NewReadWriteCounter(input io.ReadWriter) ReadWriteCounter {
	return &ReadWriteStat{
		r: input,
		w: input,
	}
}
