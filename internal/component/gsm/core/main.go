package main

import "go-short-me/internal/infrastructure/app"

func main() {
	var server app.App
	server.Start()
}
