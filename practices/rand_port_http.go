package main

import (
	"log"
	"net"
	"net/http"
)

func createListener() (l net.Listener, close func()) {
	// port 0 is a non-ephemeral port that works as a wildcard that tells the system to find
	// any available ports particularly in the Unix OS.
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	return l, func() {
		_ = l.Close()
	}
}

// It can be used when you are not sure which port is available, then use this to get an available port.
func main() {
	l, c := createListener()
	defer c()
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// handle as usual
	}))

	log.Println("listening at", l.Addr().(*net.TCPAddr).Port)
	err := http.Serve(l, nil)
	if err != nil {
		return
	}
}
