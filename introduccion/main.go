package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createURL() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic(err)
	}

	u.Host = "localhost:3000"
	u.Scheme = "http"
	query := u.Query()
	query.Add("nombre", "valor")

	u.RawQuery = query.Encode()
	return u.String()
}

func main() {
	url := createURL()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println("El header es ", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("El body es ", string(body))
	fmt.Println("El status es ", response.Status)
}
