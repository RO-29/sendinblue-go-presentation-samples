package main

import (
	"context" //https://golang.org/pkg/context/
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err, " new request")
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err, "new request")
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
