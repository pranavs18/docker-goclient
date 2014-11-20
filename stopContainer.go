/******************************************************
@author - Pranav Saxena
Rancher Labs Inc
Go-client to stop a container
Sends a post request to the Cattle server to stop a running container
****************************************************/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const containerID = "1i11"

type stopPayload struct {
	Stuff []stopData `json:"{}"`
}

type stopData struct {
	ImageUuid       string            `json:"imageUuid`
	Type            string            `json:"type`
	Id              string            `json:"id"`
	Links           Links_container   `json:"links"`
	Actions         Actions_container `json:"actions"`
	AccountID       string            `json:"accountId"`
	AgentID         string            `json:"agentId"`
	AllocationState string            `json:"allocationState"`
	Compute         string            `json:"compute"`
	Created         string            `json:"created"`
}

type stopContainer struct {
	stopcontainer []stopPayload
}

type StopContainersResponse struct {
	Data []stopContainer
}

type StopContainersOpt struct {
	stopFilters map[string]string
}

func (client *RancherClient) StopContainer(opts *StopContainersOpt) (StopContainersResponse, error) {

	url, err := url.Parse(client.Url + "/containers/" + containerID)
	url.Scheme = protocol
	url.Host = ip + separator + defaultPort
	q := url.Query()
	for k, _ := range opts.stopFilters {
		q.Set(k, opts.stopFilters[k])
	}
	url.RawQuery = q.Encode()
	fmt.Println(url.String())
	cls := &http.Client{}
	str, err1 := json.Marshal(opts.stopFilters)
	if err1 != nil {
		panic(err1)
	}
	res, _ := http.NewRequest("POST", url.String(), bytes.NewBufferString(string(str)))
	resp, _ := cls.Do(res)
	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		defer resp.Body.Close()
		fmt.Println(err)
	}

	var p stopPayload

	err = json.Unmarshal([]byte(string(body)), &p)

	if err != nil {
		fmt.Println("Error reading JSON")
		panic(err)
	}

	var result StopContainersResponse
	var temp stopContainer
	temp.stopcontainer = append(temp.stopcontainer, p)
	result.Data = append(result.Data, temp)
	fmt.Println("Fetching JSON data from the Cattle server...")

	return result, err
}
