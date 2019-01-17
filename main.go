package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Args[1]

	// create client to make requests
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 0,
	}
	// make request
	response, err := client.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/550?api_key=%s", apiKey))
	if err != nil {
		log.Panic(err)
	}

	// extract data from response
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Panic(err)
	}

	var data interface{}

	err = json.Unmarshal(responseBody, &data)

	if err != nil {
		log.Panic(err)
	}

	// print data
	fmt.Println(data)
}
