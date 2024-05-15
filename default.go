package main

import (
	"fmt"
	"net/http"
)

func defaultServer() {
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/hello/world/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World !")
	})

	fmt.Println("⇨ http server started on \033[0;32m[::]:8000\033[0m")
	http.ListenAndServe(":8000", nil)
}
