package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {

	resp, err := http.Get("http://example.com")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()	
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response:" + string(body));
	
}
