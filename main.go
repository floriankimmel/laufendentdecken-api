package main

import "log"

func main() {
	store, err := NewSQLLiteStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8000", store)
	server.Run()
}
