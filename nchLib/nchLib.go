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


func (nch *NchCred) InitNch (yamlFilNam string) (err error) {

	info, err := os.Stat(yamlFilNam)
	if err != nil {
		return fmt.Errorf("os.Stat file: %s err: %v", yamlFilNam, err)
	}

	fmt.Printf("opened file: %s\nSize %d\n", yamlFilNam, info.Size())

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

	fmt.Printf("buf: \n%s\n", string(buf))

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
