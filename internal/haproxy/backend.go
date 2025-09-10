package haproxy

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"k8s.io/apimachinery/pkg/util/json"
)

type Backend struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

func ListBackend(version Version) ([]Backend, error) {
	var haproxy = "http://localhost:5555/v3"

	url := haproxy + "/services/haproxy/configuration/backends?version=" + version.Version
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}
	req.SetBasicAuth("admin", "adminpwd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	record, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}
	var backends []Backend
	err = json.Unmarshal(record, &backends)
	if err != nil {
		log.Print(err)
	}
	return backends, nil
}

func CreateBackend(version Version, backendRequest *Backend) (Backend, error) {
	var haproxy = "http://localhost:5555/v3"

	url := haproxy + "/services/haproxy/configuration/backends?version=" + version.Version
	payload, err := json.Marshal(&backendRequest)
	if err != nil {
		log.Print(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Print(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	record, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	var backend Backend
	err = json.Unmarshal(record, &backend)
	if err != nil {
		log.Print(err)
	}

	return backend, nil
}
