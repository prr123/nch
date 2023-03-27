// test program for namecheap sandbox
// author: prr azulsoftware
// date: 26 Mar 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
    "fmt"
//    "io/ioutil"
    "net/http"
    "time"
	"log"
	"context"
)

func main() {

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("Response Status Code: %s\n", res.Status)
}

