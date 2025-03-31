package main

import (
	"fmt"
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
	fmt.Println("Starting server...")
	kvstate.InitiateKVState()

	fmt.Println("Starting API...")
	api.InitializeApi(port)
}
