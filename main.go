package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strings"
)

var db map[string]string
var letters = []rune("0123456789abcdef")
var URL_BASE = "http://localhost:8080/"

type UserRequest struct {
    Url string `json:"url"`
}

func init() {
    db = make(map[string]string)
}

func generate_string(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func hello_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func create_handler(w http.ResponseWriter, r *http.Request) {
    var ur UserRequest

    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&ur)

    short_url := generate_string(10)
    db[short_url] = ur.Url

    fmt.Fprintf(w, URL_BASE + "tiny/" + short_url)
}

func list_handler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(db)
}

func short_url_handler(w http.ResponseWriter, r *http.Request) {
    s := strings.Split(r.URL.Path, "/")
    short_path := s[2]

    long_url := db[short_path]
    fmt.Fprintf(w, long_url)
}

func main() {
    http.HandleFunc("/", hello_handler)
    http.HandleFunc("/create", create_handler)
    http.HandleFunc("/list", list_handler)
    http.HandleFunc("/tiny/", short_url_handler)

    http.ListenAndServe(":8080", nil)
}
