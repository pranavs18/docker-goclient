/******************************************************
@author - Pranav Saxena
Rancher Labs Inc
Go-client to fetch docker container's json response
****************************************************/

package main

import (
	"fmt"
	"log"
	"net/url"
)

const defaultPort = "8080"
const ip = "127.0.0.1"
const protocol = "http"
const separator = ":"
const versionAPI = "/v1"

type RancherClient struct {
	Url string
}

func NewRancherClient(url string) *RancherClient {
	return &RancherClient{
		Url: url,
	}
}

func main() {
	url, err := url.Parse(protocol + separator + "//" + ip + separator + defaultPort + versionAPI)
	if err != nil {
		log.Fatal(err)
	}

	client := NewRancherClient(url.String())
	data, errListContainer := client.ListContainers(&ListContainersOpt{
		Filters: map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
	})

	if errListContainer != nil {
		panic(errListContainer)
		log.Fatal(errListContainer)
	}

	fmt.Println("Data Retrieved .... ")
	log.Println(data.Data)

	fmt.Println("Creating new container ... ")
	client2 := NewRancherClient(url.String())
	createData, errCreateContainer := client2.CreateContainer(&CreateContainersOpt{
		createFilters: map[string]string{
			"imageUuid": "docker:nginx",
			"name":      "dummyContainer",
		},
	})

	if errCreateContainer != nil {
		panic(errCreateContainer)
		log.Fatal(errCreateContainer)
	}

	fmt.Println("Container created successfully")
	log.Println(createData.Data)

	fmt.Println("Stopping an active container ... ")
	client3 := NewRancherClient(url.String())
	stopData, errStopContainer := client3.StopContainer(&StopContainersOpt{
		stopFilters: map[string]string{
			"action": "stop",
		},
	})

	if errStopContainer != nil {
		panic(errStopContainer)
		log.Fatal(errStopContainer)
	}

	fmt.Println("Container stopped successfully")
	log.Println(stopData.Data)

}
