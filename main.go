package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/furtidev/avro.go/parsing"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	prompt := flag.String("p", "ami banglay gan gai", "Text you want to parse")
	flag.Parse()
	data, err := parsing.ConvertToJSON()
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	fmt.Println(data.Parse(prompt))
}
