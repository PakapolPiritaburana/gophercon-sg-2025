package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	return http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from goroutine %d\n", goroutineID())
		// log request async
		go func() {
			fmt.Printf("%s - - [%s] \"%s %s %s\" %d %d\n", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, r.UserAgent(), 200, 0)
		}()
	}))
}

func goroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var id int
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	return id
}
