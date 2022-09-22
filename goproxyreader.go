package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type VersionInfo struct {
	Version string
	Time    string
}

func ReadLatest(proxy string, path string) (*VersionInfo, error) {
	resp, err := http.Get("https://" + proxy + "/" + path + "/@latest")

	if err != nil {
		fmt.Printf("error while GET operation: %#v\n", err)
		return nil, err
	}

	var vi VersionInfo

	err = json.NewDecoder(resp.Body).Decode(&vi)

	return &vi, err

}

/*

   The client should at least support the info and @latest endpoints. list might be a good addition, but is not strictly required.
   mod and zip endpoints are not used in this project but can serve as an additional exercise.
   The client should support proxies running at different locations with https://proxy.golang.org/ being the default.
   All operations should receive a context.Context as the first parameter and respect any deadline applied to the context.
   Try to come up with an idiomatic API.
   Write a unit test.

*/

func ReadInfo(proxy string, path string, version string) (*VersionInfo, error) {
	resp, err := http.Get("https://" + proxy + "/" + path + "/@v/" + version + ".info")

	if err != nil {
		fmt.Printf("error while GET operation: %#v\n", err)
		return nil, err
	}

	var vi VersionInfo

	err = json.NewDecoder(resp.Body).Decode(&vi)

	return &vi, err

}

func ReadList(proxy string, path string) (*[]string, error) {
	resp, err := http.Get("https://" + proxy + "/" + path + "/@v/list")

	if err != nil {
		fmt.Printf("error while GET operation: %#v\n", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	var r = strings.Split(string(body), "\n")

	return &r, err
}

func main() {
	li, _ := ReadList("proxy.golang.org", "golang.org/x/text")
	fmt.Printf("???: %#v", *li)
}
