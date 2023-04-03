// test program for namecheap sandbox
// author: prr azulsoftware
// date: 26 Mar 2023
// copyright 2023 prr, azulsoftware
//
package main

import (
    "fmt"
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

	err := nchCred.InitNch("nchProd.yaml")
	if err != nil {
		log.Fatalf("InitNch: %v", err)
	}

	nchCred.PrintNchCred()

    clientOpt, err := nchCred.GetClientOpt()
    if err != nil {
        log.Fatalf("GetClientOpt: %v", err)
    }

    nchLib.PrintClientOptions(clientOpt)

    client := nch.NewClient(clientOpt)

	fmt.Printf("client: %v\n", client)

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

}

func PrintDomains(resp *nch.DomainsGetListCommandResponse) {

	fmt.Printf("*************** Domains [%d] *************\n", len(*resp.Domains))

	for i:=0; i< len(*resp.Domains); i++ {
		dom := (*resp.Domains)[i]
		fmt.Printf("[%d]: %s %s\n",i+1, *(dom.ID), *(dom.Name))
	}
}
