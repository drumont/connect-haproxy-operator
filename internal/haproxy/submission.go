package haproxy

import (
	"log"
)

func ReconcileIngress(backendName string) {
	version, err := GetVersion()
	if err != nil {
		log.Print(err)
	}

	log.Printf("Haproxy current version %s", version.Version)

	backends, err := ListBackend(version)
	if err != nil {
		log.Print(err)
	}

	var backend Backend

	for _, b := range backends {
		if b.Name == backendName {
			backend = b
		}
	}

	if backend == (Backend{}) {
		log.Print("No backend found. Take action to create it")
	}
}
