package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type webhook struct {
	Action     string
	Repository struct {
		ID       string
		FullName string
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":80", nil))
}

//test1
