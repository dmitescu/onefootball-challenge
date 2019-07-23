package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	baseURL string = os.Getenv("ONEFOOTBALL_BASE_URL")
)

func GetUrl(tid int) string {
	return fmt.Sprintf("%s/%d.json", baseURL, tid)
}

func QueryId(tid int) (*QueryResponse, error) {
	var err error
	var url string
	var body []byte
	
	client := &http.Client{}

	url = GetUrl(tid)
	
	request, err := http.NewRequest(
		"GET",
		url,
		nil)

	response, _ := client.Do(request)
	
	qresponse := &QueryResponse{}

	body, err = ioutil.ReadAll(response.Body)

	if (err == nil) {
		err = json.Unmarshal(body, qresponse)
	}
	
	if (err != nil || response.StatusCode != 200) {
		log.Println("status code: ", response.StatusCode)
		log.Println("error: ", err)
		return nil, err
	}

	return qresponse, nil
}

