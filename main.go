package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
