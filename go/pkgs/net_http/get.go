package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// *http.Response, error
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s", robots)
}
