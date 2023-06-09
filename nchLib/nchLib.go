// namecheap library
// author: prr, azul software
// date: 27 Mar 2022
// copyright 2023 prr, azul software
//
package nchLib

import (
	"fmt"
	"os"
	yaml "github.com/goccy/go-yaml"
    nchSdk "ns/nchReg/nchSdk"
	)

type NchCred struct {
	Sandbox bool `yaml:"Sandbox"`
	Email string `yaml:"Email"`
	ApiUser string `yaml:"ApiUser"`
	ApiKey string `yaml:"ApiKey"`
	IP string `yaml:"IP"`
	Domain string `yaml:"Domain"`
	File string
	NchApi string
}

type CfNsNames struct {
	File string
	Domain string `yaml:"domain"`
	Ns1 string `yaml:"ns1"`
	Ns2 string `yaml:"ns2"`
}

func (nch *NchCred) InitNch (yamlFilNam string, dbg bool) (err error) {

	info, err := os.Stat(yamlFilNam)
	if err != nil {
		return fmt.Errorf("os.Stat file: %s err: %v", yamlFilNam, err)
	}

	if dbg {fmt.Printf("opened file: %s\nSize %d\n", yamlFilNam, info.Size())}

	infil, err := os.Open(yamlFilNam)
	if err != nil {
		return fmt.Errorf("os.Open file: %s err: %v", yamlFilNam, err)
	}

	nch.File = yamlFilNam
	buf := make([]byte, int(info.Size()))

	_, err = infil.Read(buf)
	if err != nil {
		return fmt.Errorf("Read err: %v", err)
	}

	if dbg {fmt.Printf("buf: \n%s\n", string(buf))}

    if err := yaml.Unmarshal(buf, nch); err !=nil {
        return fmt.Errorf("error Unmarshall: %v\n", err)
    }

	baseUrl := `https://api.namecheap.com/xmlresponse?`
	if nch.Sandbox {baseUrl = `https://api.sandbox.namecheap.com/xmlresponse?`}

	userStr := "ApiUser=" + nch.ApiUser
	userNamStr := "&UserName=" + nch.ApiUser
	apiKeyStr := "&ApiKey=" + nch.ApiKey
	ipStr := "&ClientIp=" + nch.IP
	cmdStr :="&Command=xxx"
	domainStr := "&DomainList=" + nch.Domain
	baseUrl += userStr + apiKeyStr + userNamStr + ipStr + cmdStr + domainStr
	nch.NchApi = baseUrl
	return nil
}

func (nch *NchCred) GetClientOpt()  (clOptRef *nchSdk.ClientOptions, err error) {
	var clOpt nchSdk.ClientOptions

	clOpt.UserName = nch.ApiUser
	clOpt.ApiUser = nch.ApiUser
	clOpt.ApiKey = nch.ApiKey
	clOpt.ClientIp = nch.IP
	clOpt.UseSandbox = nch.Sandbox

	return &clOpt, nil
}

func ReadNsRec(yamlFilNam string, dbg bool) (ns *CfNsNames, err error){

	var nsRec CfNsNames

    infil, err := os.Open(yamlFilNam)
    if err != nil {
        return nil, fmt.Errorf("os.Open file: %s err: %v", yamlFilNam, err)
    }

    nsRec.File = yamlFilNam
    buf := make([]byte, 250)

    _, err = infil.Read(buf)
    if err != nil {
        return nil, fmt.Errorf("Read err: %v", err)
    }

    if dbg {fmt.Printf("buf: \n%s\n", string(buf))}

    if err := yaml.Unmarshal(buf, &nsRec); err !=nil {
        return nil, fmt.Errorf("error Unmarshall: %v\n", err)
    }

	if dbg {
		fmt.Printf("File:   %s\n", nsRec.File)
		fmt.Printf("Domain: %s\n", nsRec.Domain)
		fmt.Printf("ns1:  %s\n", nsRec.Ns1)
		fmt.Printf("ns2:  %s\n", nsRec.Ns2)
	}
	return &nsRec, nil
}

func (nch *NchCred) PrintNchCred () {

    fmt.Println("**** NCH Credentials *****")
    fmt.Printf("NCH User:  %s\n", nch.ApiUser)
    fmt.Printf("NCH Key:   %s\n", nch.ApiKey)
    fmt.Printf("Email:     %s\n", nch.Email)
    fmt.Printf("Sandbox:   %t\n", nch.Sandbox)
	fmt.Printf("IP:        %s\n", nch.IP)
	fmt.Printf("Domain:    %s\n", nch.Domain)
	fmt.Printf("File:      %s\n", nch.File)
	fmt.Printf("NCP Api:   %s\n", nch.NchApi)
    fmt.Println("**************************")
}

func PrintClientOptions (nch *nchSdk.ClientOptions) {

    fmt.Println("**** Client Options *****")
    fmt.Printf("UserName:  %s\n", nch.UserName)
    fmt.Printf("ApiUser:   %s\n", nch.ApiUser)
    fmt.Printf("ApiKey:    %s\n", nch.ApiKey)
	fmt.Printf("ClientIp:  %s\n", nch.ClientIp)
    fmt.Printf("UseSandbox:%t\n", nch.UseSandbox)
    fmt.Println("**************************")

}

func PrintDomains(resp *nchSdk.DomainsGetListCommandResponse) {

    fmt.Printf("*************** Domains [%d] *************\n", len(*resp.Domains))

    for i:=0; i< len(*resp.Domains); i++ {
        dom := (*resp.Domains)[i]
        fmt.Printf("[%d]: %s %s\n",i+1, *(dom.ID), *(dom.Name))
    }
}

