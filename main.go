package main

import (
	"fmt"
	"net/http"

	"github.com/Atoo35/http-mock-test/http_client"
)

func main() {
	TestCall()
}

func TestCall() {
	// http_client.Client.New()

	client := http_client.Client

	req, _ := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("resp", resp)
}
