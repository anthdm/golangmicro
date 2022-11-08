package main

import "flag"

func main() {
	port := flag.String("port", ":3000", "listen address of the service")
	flag.Parse()

	svc := loggingService{priceService{}}
	server := NewJSONAPIServer(*port, svc)
	server.Run()
}
