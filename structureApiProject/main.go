package main

import (
	"GoExpGain/structureApiProject/api"
	"GoExpGain/structureApiProject/storage"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "my server address")
	flag.Parse()

	mystore := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, mystore)
	fmt.Println("Server running on port: ", *listenAddr)
	log.Fatal(server.Start())

}
