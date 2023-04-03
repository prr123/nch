// setNsCf.go
// program that sets the nameservers to cloudflare nameservers
//
// author: prr azulsoftware
// date: 26 Mar 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
//    "fmt"
//    "io/ioutil"
//    "net/http"
//    "time"
	"log"
//	"context"
		"ns/nchReg/nchLib"
    nch "ns/nchReg/nchSdk"
)



func main() {

	var nchCred nchLib.NchCred
	// read yaml file for user, ip and key

	err := nchCred.InitNch("nchProd.yaml", true)
	if err != nil {
		log.Fatalf("InitNch: %v", err)
	}

//	nchCred.PrintNchCred()

    clientOpt, err := nchCred.GetClientOpt()
    if err != nil {
        log.Fatalf("GetClientOpt: %v", err)
    }

//    nchLib.PrintClientOptions(clientOpt)

	nsRecFilNam := "nsCf.Yaml"
	nsRec, err := nchLib.ReadNsRec(nsRecFilNam ,true)
	if err != nil {
        log.Fatalf("ReadNsRec: %v", err)
	}

	log.Printf("ns1: %s\n", nsRec.Ns1)
	log.Printf("ns2: %s\n", nsRec.Ns2)
    client := nch.NewClient(clientOpt)

	log.Printf("client: %v\n", client)

/*
	var nsListAr [2]string
	nsListAr[0] := nsRec.Ns1
	nsListAr[0] := nsRec.Ns2

	nsList := nsListAr[:]

	resp, err := client.DomainsDNS.SetCustom(nsRec.Domain, nsList)
	if err!= nil {
		log.Fatalf("DomainDns.SetCustom: %v", err)
	}


	listStr := "ALL"
	pageSize := 100
	sortStr := "NAME"
	domainParams := &nch.DomainsGetListArgs{
		ListType: &listStr,
		PageSize: &pageSize,
		SortBy: &sortStr,
	}

	domainListResp, err := client.Domains.GetList(domainParams)
	if err!= nil {
		log.Fatalf("GETLIST: %v", err)
	}

	PrintDomains(domainListResp)
*/
}

