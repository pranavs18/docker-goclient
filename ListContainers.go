/******************************************************
@author - Pranav Saxena
Rancher Labs Inc
Go-client to fetch docker container's json response
****************************************************/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Payload struct {
	Stuff []Data `json:"data"`
}

type Data struct {
	Id              string            `json:"id"`
	Links           Links_container   `json:"links"`
	Actions         Actions_container `json:"actions"`
	AccountID       string            `json:"accountId"`
	AgentID         string            `json:"agentId"`
	AllocationState string            `json:"allocationState"`
	Compute         string            `json:"compute"`
	Created         string            `json:"created"`
}

type Links_container map[string]string
type Actions_container map[string]string

type Container struct {
	container []Payload
}

type ListContainersResponse struct {
	Data []Container
}

type ListContainersOpt struct {
	Filters map[string]string
}

// Go function to retrieve container's response
func (client *RancherClient) ListContainers(opts *ListContainersOpt) (ListContainersResponse, error) {
	// fetch the base URL
	url, err := url.Parse(client.Url + "/containers")
	url.Scheme = protocol
	url.Host = ip + separator + defaultPort
	q := url.Query()
	for k, _ := range opts.Filters {
		q.Set(k, opts.Filters[k])
	}
	url.RawQuery = q.Encode()
	res, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		defer res.Body.Close()
		fmt.Println(err)
	}

	var p Payload

	err = json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	var result ListContainersResponse
	var temp Container
	temp.container = append(temp.container, p)
	result.Data = append(result.Data, temp)
	fmt.Println("Fetching JSON data from the Cattle server...")
	for _, stuff := range result.Data {
		fmt.Println(" Account ID "+stuff.container[0].Stuff[0].AccountID, "\n",
			stuff.container[0].Stuff[0].Actions, "\n",
			"Agent ID "+stuff.container[0].Stuff[0].AgentID, "\n",
			"Allocation State "+stuff.container[0].Stuff[0].AllocationState, "\n",
			"Compute "+stuff.container[0].Stuff[0].Compute, "\n",
			"Created "+stuff.container[0].Stuff[0].Created, "\n",
			"ID "+stuff.container[0].Stuff[0].Id, "\n",
			stuff.container[0].Stuff[0].Links)

	}
	return result, err
}
