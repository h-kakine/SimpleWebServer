package main

import (
        "io"
	"fmt"
	"log"
        "encoding/json"
	"net/http"
)

func defaultHead(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        log.Printf("path: %v", path)
        query := r.URL.Query()
        log.Printf("QueryString: %v", query)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNoContent)
        return
}

func defaultGet(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        log.Printf("path: %v", path)
        query := r.URL.Query()
        log.Printf("QueryString: %v", query)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
        m := map[string]string{"status": "ok"}
        if err := json.NewEncoder(w).Encode(m); err != nil {
            log.Println(err)
        }
        return
}

func defaultPost(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        log.Printf("path: %v", path)
        query := r.URL.Query()
        log.Printf("QueryString: %v", query)
        content := r.Header.Get("Content-Type")
        log.Printf("Content-Type: %v", content) 
        if (content == "application/x-www-form-urlencoded") {
            b, _ := io.ReadAll(r.Body)
            log.Println(string(b))
        } else if (content == "application/json") {
            b, _ := io.ReadAll(r.Body)
            log.Println(string(b))
        }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
        m := map[string]string{"status": "ok"}
        if err := json.NewEncoder(w).Encode(m); err != nil {
            log.Println(err)
        }
}

func defaultPut(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        log.Printf("path: %v", path)
        query := r.URL.Query()
        log.Printf("QueryString: %v", query)
        content := r.Header.Get("Content-Type")
        log.Printf("Content-Type: %v", content) 
        if (content == "application/x-www-form-urlencoded") {
            b, _ := io.ReadAll(r.Body)
            log.Println(b)
        } else if (content == "application/json") {
            b, _ := io.ReadAll(r.Body)
            log.Println(b)
        }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
        m := map[string]string{"status": "ok"}
        if err := json.NewEncoder(w).Encode(m); err != nil {
            log.Println(err)
        }
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
        method := r.Method
        log.Printf("method: %v", method)
        switch method {
       	case "HEAD":
		defaultHead(w,r)
        case "GET":
		defaultGet(w,r)
        case "POST":
		defaultPost(w,r)
        case "PUT":
		defaultPut(w,r)
	}
	return
}

func main() {
	http.HandleFunc("/", defaultHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
