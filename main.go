package main

import "github.com/marcelofelixsalgado/financial-TEMPLATE-api/api"

func main() {
	// Start HTTP Server
	server := api.NewServer()
	server.Run()
}
