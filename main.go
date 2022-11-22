package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, error := os.Hostname()

		if error != nil {
			fmt.Println(error)
			w.Write([]byte(string(error.Error())))
		}

		w.Write([]byte(fmt.Sprintf("<h1>Hostname: %s</h1>", hostname)))
	})
	http.ListenAndServe(":8080", nil)
}
