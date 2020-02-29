package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// const timeout = time.Second * 10
const timeout = time.Millisecond * 10
const url = "http://google.com.vn"

func main() {
	log.Println("begin main...")
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, timeout)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Request failed:", err)
		return
	}
	log.Println("Resp with status code:", res.StatusCode)
}
