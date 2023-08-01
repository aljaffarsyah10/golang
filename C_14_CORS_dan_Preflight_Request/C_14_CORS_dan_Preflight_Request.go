package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
        // w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		// Multiple Origin
		w.Header().Set("Access-Control-Allow-Origin", 
        "https://www.google.com, https://novalagung.com")
		// w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
        // w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		// semua origin mendapat ijin akses
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// semua method diperbolehkan masuk
		w.Header().Set("Access-Control-Allow-Methods", "*")

		// semua header diperbolehkan untuk disisipkan
		w.Header().Set("Access-Control-Allow-Headers", "*")		

        if r.Method == "OPTIONS" {
            w.Write([]byte("allowed"))
            return
        }

        w.Write([]byte("hello"))
    })

    log.Println("Starting app at :9000")
    http.ListenAndServe(":9000", nil)
}