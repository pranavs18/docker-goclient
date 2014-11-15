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
	"net/http"
)

const defaultPort = "8080"
const ip = "127.0.0.1"

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

/*func serveRest(w http.ResponseWriter, r *http.Request) {
	response := getJsonResponse()
}*/

func main() {
	url := "http://" + ip + ":" + defaultPort + "/v1/containers"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
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

	for _, stuff := range p.Stuff {
		fmt.Println(stuff.AccountID, "\n", stuff.Actions, "\n",
			stuff.AgentID, "\n", stuff.AllocationState, "\n", stuff.Compute,
			"\n", stuff.Created, "\n", stuff.Id, "\n", stuff.Links)
	}
}
