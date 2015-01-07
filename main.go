// Package main provides webserver for smatercar.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8000, "port to serve on")
	flag.Parse()

	webFileHanlder := http.FileServer(http.Dir("dist/"))
	musicFileHandler := http.FileServer(http.Dir(""))

	http.Handle("/", webFileHanlder)
	http.Handle("/music/", musicFileHandler)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf(":%d", *port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
