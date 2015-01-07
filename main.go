// Package main provides webserver for smatercar.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bjorngi/go-carweb/media"
	"log"
	"net/http"
)

//func getMusicTracks(trackChan chan *[]media.Track) {
//	tracks, err := media.GetTracks()
//	if err != nil {
//		panic("Failed to get tracks")
//	}
//	time.Sleep(time.Second * 2)
//
//	trackChan <- tracks
//}

func musicHandleFunction(res http.ResponseWriter, req *http.Request) {
	trackChan := make(chan *[]media.Track)

	go func() {
		tracks, err := media.GetTracks("music")
		if err != nil {
			fmt.Printf("Could not get tracks: %v\n", err)
		}
		trackChan <- tracks

	}()

	tracks := <-trackChan

	b, err := json.Marshal(tracks)
	if err != nil {
		fmt.Printf("Failed to encode json\n")
	}

	res.Write(b)

}

func main() {
	port := flag.Int("port", 8000, "port to serve on")
	flag.Parse()

	webFileHanlder := http.FileServer(http.Dir("dist/"))
	musicFileHandler := http.FileServer(http.Dir(""))

	http.Handle("/", webFileHanlder)
	http.Handle("/music/", musicFileHandler)
	http.HandleFunc("/music/list", musicHandleFunction)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf(":%d", *port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
