package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("error while GET operation: %#v\n", err)
		return
	}

	fmt.Printf("status:      %s\n", resp.Status)
	fmt.Printf("status code: %d\n", resp.StatusCode)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error while reading BODY: %#v\n", err)
		return
	}

	fmt.Printf("body:        %s\n", body)
}
