package main

//go:generate go run generators/main.go

import (
	"fincal-server/server"
	"flag"
	_ "github.com/bufbuild/protovalidate-go"
	_ "github.com/joho/godotenv/autoload"
)

var (
	migrate = flag.String("migrate", "", "Migrate DB")
)

func main() {
	flag.Parse()

	if *migrate == "up" {
		migrateUp()
		return
	}

	if *migrate == "down" {
		migrateDown()
		return
	}

	server.InitServer()
}
