package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/public/")
		w.WriteHeader(http.StatusPermanentRedirect) // 308
	})

	http.HandleFunc("/301", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMovedPermanently)
		fmt.Fprintf(w, templateHeader, http.StatusMovedPermanently)
	})

	http.HandleFunc("/302", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, templateHeader, http.StatusFound)
	})

	http.HandleFunc("/303", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusSeeOther)
		fmt.Fprintf(w, templateHeader, http.StatusSeeOther)
	})

	http.HandleFunc("/307", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprintf(w, templateHeader, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/308", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPermanentRedirect)
		fmt.Fprintf(w, templateHeader, http.StatusPermanentRedirect)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Println("PORT:", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

var templateHeader = `<!DOCTYPE html>
<html>
<head>
    <title>header landing</title>
</head>
<body>
    <h1>%v header landing</h1>
    <p><a href="/">home</a></p>
</body>
</html>`
