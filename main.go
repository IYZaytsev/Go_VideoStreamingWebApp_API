package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "randomVideo.mp4")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
