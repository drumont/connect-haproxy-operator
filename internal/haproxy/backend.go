package haproxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Backend struct {
	Name string
}

func ListBackends() {
	var haproxy string = ""
	url := haproxy + "/services/haproxy/configuration/backends"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	record, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(record)

}
