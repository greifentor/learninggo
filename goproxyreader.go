package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LatestInfo struct {
	Version string
	Time    string
}

func ReadLatest(proxy string, path string) (*LatestInfo, error) {
	resp, err := http.Get("https://" + proxy + "/" + path + "/@latest")

	if err != nil {
		fmt.Printf("error while GET operation: %#v\n", err)
		return nil, err
	}

	var li LatestInfo

	err = json.NewDecoder(resp.Body).Decode(&li)

	return &li, err

}

/*

    The client should at least support the info and @latest endpoints. list might be a good addition, but is not strictly required.
    mod and zip endpoints are not used in this project but can serve as an additional exercise.
    The client should support proxies running at different locations with https://proxy.golang.org/ being the default.
    All operations should receive a context.Context as the first parameter and respect any deadline applied to the context.
    Try to come up with an idiomatic API.
    Write a unit test.

*/

func ReadInfo() {

}

func ()
