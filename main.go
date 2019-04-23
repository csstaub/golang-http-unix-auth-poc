package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	addr, err := net.ResolveUnixAddr("unix", "/tmp/unix.sock")
	panicOnError(err)

	listener, err := net.ListenUnix("unix", addr)
	panicOnError(err)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", r.RemoteAddr)
	})

	listener.SetUnlinkOnClose(true)
	panicOnError(http.Serve(&authenticatedListener{listener}, mux))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
