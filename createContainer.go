/******************************************************
@author - Pranav Saxena
Rancher Labs Inc
Go-client to create a container
Sends a post request to the Cattle server with imageUuid
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

type createPayload struct {
	Stuff []createData `json:"{}"`
}

type createData struct {
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

type createContainer struct {
	createcontainer []createPayload
}

type CreateContainersResponse struct {
	Data []createContainer
}

type CreateContainersOpt struct {
	createFilters map[string]string
}

func (client *RancherClient) CreateContainer(opts *CreateContainersOpt) (CreateContainersResponse, error) {

	url, err := url.Parse(client.Url + "/containers")
	url.Scheme = protocol
	url.Host = ip + separator + defaultPort
	q := url.Query()
	for k, _ := range opts.createFilters {
		q.Set(k, opts.createFilters[k])
	}
	url.RawQuery = q.Encode()
	fmt.Println(url.String())
	cls := &http.Client{}
	str, err1 := json.Marshal(opts.createFilters)
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

	var p createPayload

	err = json.Unmarshal([]byte(string(body)), &p)

	if err != nil {
		fmt.Println("Error reading JSON")
		panic(err)
	}

	var result CreateContainersResponse
	var temp createContainer
	temp.createcontainer = append(temp.createcontainer, p)
	result.Data = append(result.Data, temp)
	fmt.Println("Fetching JSON data from the Cattle server...")

	return result, err
}
