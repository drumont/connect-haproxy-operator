package haproxy

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type Version struct {
	Version string
}

func GetVersion() (Version, error) {
	var haproxy string = "http://localhost:5555/v3"
	url := haproxy + "/services/haproxy/configuration/version"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}

	req.SetBasicAuth("admin", "adminpwd")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	version, err := io.ReadAll(resp.Body)
	v := strings.TrimSpace(string(version))
	return Version{v}, nil
}
