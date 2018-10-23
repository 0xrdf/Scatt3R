package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func ShowName() {
	fmt.Println("[*] Scatt3R started ! [*]")
}

func ServeFiles() {
	ShowName()
	port := flag.String("p", "1337", "Serve")
	directory := flag.String("d", ".", "scatter Directory")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("\n\nServing directory %s on HTTP port: %s\n", *directory, *port)
	fmt.Println("\nClose this window to stop.")
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
