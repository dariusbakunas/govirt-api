package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"log"
	"geekspace.us/govirt-api/server"
	"geekspace.us/govirt-api/models"
)

func main() {
	parser := argparse.NewParser("govirt-api", "Libvirt Rest API")

	uri := parser.String("u", "uri", &argparse.Options{Required: true, Help: "Libvirt URI"})
	port := parser.Int("p", "port", &argparse.Options{Required: false, Help: "Listening port"})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	} else {
		if *port == 0 {
			*port = 8000
		}

		conn, err := models.InitLibvirt(*uri)
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}

		defer conn.Close()

		server.Init(*port)
	}
}