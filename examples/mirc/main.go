package main

import (
	"log"

	. "github.com/alimy/mir/v3/core"
	. "github.com/alimy/mir/v3/engine"

	_ "github.com/alimy/mir-example/v3/mirc/routes"
	_ "github.com/alimy/mir-example/v3/mirc/routes/v1"
	_ "github.com/alimy/mir-example/v3/mirc/routes/v2"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialMode),
		GeneratorName(GeneratorGin),
		SinkPath("auto"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
