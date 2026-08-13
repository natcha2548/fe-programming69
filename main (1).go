package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "สวัสดีจากครัวของผมเอง!",
        })
    })
    http.ListenAndServe(":8080", nil)
}