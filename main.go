package main

import (
	"log"
	"os"

	"key_val_store/api"
	kvstate "key_val_store/kv_state"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a port.")
	}

	port := os.Args[1]

	// Initialize state
	kvstate.InitiateKVState()

	api.InitializeApi(port)
}
