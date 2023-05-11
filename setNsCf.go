// setNsCf.go
// program that sets the nameservers to cloudflare nameservers
//
// author: prr azulsoftware
// date: 26 Mar 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
	"log"
	"os"
    "fmt"
//    "io/ioutil"
//    "net/http"
//    "time"
//	"context"
		"ns/nchReg/nchLib"
    nch "ns/nchReg/nchSdk"
)


func main() {

	var nchCred nchLib.NchCred

	numArg := len(os.Args)

	nsRecFilNam := "nsCf.Yaml"
	apiFilNam := "nchProd.yaml"

	useStr := "usage: ./setNsCf domain\ndomain: yaml file\n"
	switch numArg {
	case 1:
	case 2:
		nsRecFilNam = os.Args[1]
	default:
		fmt.Printf("too many cli args:\n")
		fmt.Printf(useStr)
		os.Exit(-1)
	}

	// read yaml file for user, ip and key
	log.Printf("using api: %s\n", apiFilNam)
	log.Printf("using ns: %s\n", nsRecFilNam)

	err := nchCred.InitNch(apiFilNam, true)
	if err != nil {
		log.Fatalf("InitNch: %v", err)
	}

//	nchCred.PrintNchCred()

    clientOpt, err := nchCred.GetClientOpt()
    if err != nil {
        log.Fatalf("GetClientOpt: %v", err)
    }

//    nchLib.PrintClientOptions(clientOpt)

	nsRec, err := nchLib.ReadNsRec(nsRecFilNam ,true)
	if err != nil {
        log.Fatalf("ReadNsRec: %v", err)
	}

	log.Printf("ns1: %s\n", nsRec.Ns1)
	log.Printf("ns2: %s\n", nsRec.Ns2)

    client := nch.NewClient(clientOpt)

	log.Printf("client: %v\n", client)


	var nsListAr [2]string
	nsListAr[0] = nsRec.Ns1
	nsListAr[1] = nsRec.Ns2

	nsList := nsListAr[:]

	resp, err := client.DomainsDNS.SetCustom(nsRec.Domain, nsList)
	if err!= nil {
		log.Fatalf("DomainDns.SetCustom: %v", err)
	}

	res := resp.DomainDNSSetCustomResult

	log.Printf("domain: %s\n", *(res.Domain))
	log.Printf("Updated: %t\n", *(res.Updated))

}

