package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	return http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from goroutine %d\n", goroutineID())
		go logRequest(*r)
	}))
}

func logRequest(r http.Request) {
	fmt.Fprintf(os.Stdout, "%s - - [%s] \"%s %s %s\" %d %d\n", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, r.UserAgent(), 200, 0)
}

func goroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	var id int
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	return id
}
