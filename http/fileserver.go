package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/cgi/printenv", printenv)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Fprintln(os.Stderr, "err:", err.Error())
		os.Exit(1)
	}
}

func printenv(w http.ResponseWriter, r *http.Request) {
	envs := os.Environ()
	w.Write([]byte("<h1>Env variables:</h1><pre>"))
	for _, e := range envs {
		w.Write([]byte(e + "\n"))
	}
	w.Write([]byte("</pre>"))
}
