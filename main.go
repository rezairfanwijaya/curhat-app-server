package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("endpoint / hit with status 200OK")
		response := struct {
			Status  string `json:"status"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "hello world",
		}

		res, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	domain := fmt.Sprintf(":%v", port)

	log.Println("server running on", domain)
	if err := http.ListenAndServe(domain, nil); err != nil {
		log.Fatal(err)
	}
}
