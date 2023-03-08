package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

// Pool makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.
// An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused
// by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.
// For example, a free list maintained as part of a short-lived object is not a suitable use for a Pool,

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer types,
		// since a pointer can be put into the return interface value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Record(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	Record(os.Stdout, "path", "/search?q=flowers")
}
