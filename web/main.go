package main

import (
    "log"
    "net/http"
    "path/filepath"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    http.ServeFile(w, r, filepath.Join("static", "errors", "404.html"))
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusBadRequest)
    http.ServeFile(w, r, filepath.Join("static", "errors", "400.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        badRequestHandler(w, r)
        return
    }
    if r.URL.Path != "/" {
        notFoundHandler(w, r)
        return
    }
    http.ServeFile(w, r, filepath.Join("static", "index.html"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        badRequestHandler(w, r)
        return
    }
    http.ServeFile(w, r, filepath.Join("static", "about.html"))
}

func factsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        badRequestHandler(w, r)
        return
    }
    http.ServeFile(w, r, filepath.Join("static", "facts.html"))
}

func learnMoreHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        badRequestHandler(w, r)
        return
    }
    http.ServeFile(w, r, filepath.Join("static", "learn-more.html"))
}

func main() {
    // Serve static files (CSS)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Register page handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/facts", factsHandler)
    http.HandleFunc("/learn-more", learnMoreHandler)

    // Start the server
    log.Println("Server starting on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

/*
Tasks:
-  Create a web server using the net/http package that:
   - Serves static files (HTML, CSS) from the 'static' directory
   - Handles multiple routes: /, /about, /facts, /learn-more
   - Only accepts GET requests (returns 400 Bad Request for other methods)
   - Implements proper error handling (404 for Not Found)
   - Included in the static folder is the html and css files you will be using
*/
