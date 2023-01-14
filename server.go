package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	version := os.Getenv("VERSION")
	appName := os.Getenv("APP_NAME")

	w.Write([]byte("<h1>This is a test server!</h1>"))
	fmt.Fprintf(w, "<p><b>System version:</b> %s</p>", version)
	fmt.Fprintf(w, "<p><b>App Name:</b> %s</p>", appName)
}
