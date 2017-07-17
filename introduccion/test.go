package main

import (
	"fmt"
	"log"
	"net/http"
)

// func createURL() string {
// 	u, err := url.Parse("/params")
// 	if err != nil {
// 		panic(err)
// 	}

// 	u.Host = "localhost:3000"
// 	u.Scheme = "http"
// 	query := u.Query()
// 	query.Add("nombre", "valor")

// 	u.RawQuery = query.Encode()
// 	return u.String()
// }

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo.")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
