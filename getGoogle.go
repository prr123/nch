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
)

func main() {
    c := http.Client{Timeout: time.Duration(1) * time.Second}
    resp, err := c.Get("https://www.google.com")
    if err != nil {
        fmt.Printf("Error %s", err)
        return
    }
    defer resp.Body.Close()

	fmt.Printf("resp code: %s\n", resp.Status)

//    body, err := ioutil.ReadAll(resp.Body)
//    fmt.Printf("Body : %s", body)
}

