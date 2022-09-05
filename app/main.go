package main

import (
	"flag"

	"github.com/jhonatanlteodoro/verify_test/app/initalization"
)

func main() {
	a := initalization.App{}
	a.Initilize()

	host := flag.String("host", "127.0.0.1", "Host to run this server")
	port := flag.String("port", "8080", "Port to run this server")
	flag.Parse()

	a.Run(*host, *port)
}
