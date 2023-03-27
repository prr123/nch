// test program for namecheap sandbox
// author: prr azulsoftware
// date: 26 Mar 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
	"log"
	"context"
	"ns/namecheap/nchLib"
)



func main() {

	var nch nchLib.NchCred
	// read yaml file for user, ip and key

	err := nch.InitNch("nch.yaml")
	if err != nil {
		log.Fatalf("InitNch: %v", err)
	}

	nch.PrintNchCred()

	getStr := "https://api.sandbox.namecheap.com/xml.response?ApiUser=azulTest&ApiKey=0e5d605ad8cf45fc8fa45e4c529c0448&UserName=azulTest&ClientIp=89.116.30.49&Command=namecheap.domains.check&DomainList=azulTest.com"
	req, err := http.NewRequest("GET", getStr, nil)
	if err != nil {
		log.Fatalf("Req: %v", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Do: %v", err)
	}

	fmt.Printf("Response Status Code: %s\n", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read: %v", err)
	}

    fmt.Printf("Body:\n %s\n", body)

}

