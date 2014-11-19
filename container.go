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

const defaultPort = "8080"
const ip = "127.0.0.1"

type RancherClient struct {
	Url string
}

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

func NewRancherClient(url string) *RancherClient {
	return &RancherClient{
		Url: url,
	}
}

type Container struct {
	container []Payload
}

type ListContainersResponse struct {
	Data []Container
}

type ListContainersOpt struct {
	Filters map[string]string
}

/*func (client *RancherClient) ListContainers() (ListContainersResponse, err) {

}*/

func (client *RancherClient) ListContainers(opts *ListContainersOpt) (ListContainersResponse, error) {
	url, err := url.Parse(client.Url + "/containers")
	url.Scheme = "http"
	url.Host = ip + ":" + defaultPort
	q := url.Query()
	for k, _ := range opts.Filters {
		q.Set(k, opts.Filters[k])
	}
	url.RawQuery = q.Encode()
	fmt.Println(url.String())
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

	/*for _, stuff := range p.Stuff {
		fmt.Println(stuff.AccountID, "\n", stuff.Actions, "\n",
			stuff.AgentID, "\n", stuff.AllocationState, "\n", stuff.Compute,
			"\n", stuff.Created, "\n", stuff.Id, "\n", stuff.Links)
	}*/

	var result ListContainersResponse
	var temp Container
	temp.container = append(temp.container, p)
	result.Data = append(result.Data, temp)
	fmt.Println("Fetching JSON data from Cattle server...")
	fmt.Println(result.Data)
	/*for _, stuff := range p.Stuff {
		result.Data = append(result.Data, )
	}
	*/
	return result, err
}

func main() {
	url, err := url.Parse("http://" + ip + ":" + defaultPort + "/v1")
	if err != nil {
		log.Fatal(err)
	}

	client := NewRancherClient(url.String())
	data, err2 := client.ListContainers(&ListContainersOpt{
		Filters: map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
	})

	if err2 != nil {
		panic(err2)
	}
	/*
		for _, stuff := range data.Data {

		}*/
	fmt.Print("JSON data retrieved", data.Data)
}
