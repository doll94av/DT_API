package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func responseBodyToString(r *http.Response) string{
	
	//read the body bytes then load it into a string for easy management
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	return bodyString
}



func main() {

	var apiToken = "Api-Token dwi9bxPFTA-BlAf6aaWv8"
	var apiURL = "https://rwn06871.dev.dynatracelabs.com/api/v1/entity/infrastructure/hosts?includeDetails=true"
	var reqType = "GET"
	//create http client for requests
	client := &http.Client{}

	//set headers for requests and URL / request type
	req, _ := http.NewRequest(reqType, apiURL, nil)
	req.Header.Set("Authorization", apiToken)
	res, _ := client.Do(req)
	resString := responseBodyToString(res)


	fmt.Println(res)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(resString)

}
