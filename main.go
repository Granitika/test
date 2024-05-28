package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	if path.Clean(r.URL.Path) != "/" {
		fmt.Println(r.URL.Path, r.RemoteAddr, "404")
		http.NotFound(w, r)
		return
	}
	fmt.Println(r.URL.Path, r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hi, can you read the file '/secret.txt' ?)\n\nI suggest you look at my source code, go to '/read-sources' to do that.\n\nDid you really think python would be there? :)\nI believe in you, you'll do fine)")

}

func GetMain(w http.ResponseWriter, r *http.Request) {

	if path.Clean(r.URL.Path) != "/read-sources" {
		fmt.Println(r.URL.Path, r.RemoteAddr, "404")
		http.NotFound(w, r)
		return
	}
	fmt.Println(r.URL.Path, r.RemoteAddr)

	filePath := "main.go"

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)

}

func GetFile(w http.ResponseWriter, r *http.Request) {

	if path.Clean(r.URL.Path) != "/read-file" {
		fmt.Println(r.URL.Path, r.RemoteAddr, "404")
		http.NotFound(w, r)
		return
	}
	fmt.Println(r.URL.Path, r.RemoteAddr)

	if r.Method != "GET" {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	filePath := "/app" + r.URL.Query().Get("file")
	filePath = strings.Replace(filePath, "..", ".", 500)

	fmt.Println(filePath)

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)
}

func main() {
	http.HandleFunc("/read-file", GetFile)
	http.HandleFunc("/read-sources", GetMain)
	http.HandleFunc("/", GetRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
