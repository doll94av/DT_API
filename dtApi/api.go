package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type callSetup struct {
	url string
	method string
}

func responseBodyToString(r *http.Response) string {

	//read the body bytes then load it into a string for easy management
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	return bodyString
}

//we want to call the client with the desired API call and return responseBodyToString with the resposne
//Create a http client, then do the request passed in the args, send the response to the toString and return
func callApi(req *http.Request) string {
	client := &http.Client{}
	res, _ := client.Do(req)
	return responseBodyToString(res)
}

func main() {

	//default values for the API
	var apiToken = "Api-Token dwi9bxPFTA-BlAf6aaWv8"
	var apiURL = "https://rwn06871.dev.dynatracelabs.com/api/v1/entity/infrastructure/hosts?includeDetails=true"
	var reqType = "GET"

	//stdin reader
	reader := bufio.NewReader(os.Stdin)

	//set headers for requests and URL / request type
	req, _ := http.NewRequest(reqType, apiURL, nil)
	req.Header.Set("Authorization", apiToken)

	responseString := callApi(req)
	fmt.Println(responseString)

	fmt.Println("Press 1 for Hosts, infrastructure or 2 for cluster time")
	text, _ := reader.ReadString('\n')
	
	//idk it was getting mad at me for not declaring it
	var callsetup = callSetup{"NIL", "NIL"}

	//switch for cases (api calls), eventually this will read in from a frontend instead of cmdline
	switch call := text; call {
	case "1\n":
	callsetup = callSetup{"https://rwn06871.dev.dynatracelabs.com/api/v1/entity/infrastructure/hosts?includeDetails=true", "GET"}
	case "2\n":
	callsetup = callSetup{"https://rwn06871.dev.dynatracelabs.com/api/v1/time", "GET"}
	}

	req, _ = http.NewRequest(callsetup.method, callsetup.url, nil)
	req.Header.Set("Authorization", apiToken)
	responseString = callApi(req)
	fmt.Println(responseString)
}
