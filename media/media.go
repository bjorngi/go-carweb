// Package media provides support for listing and identifying music on external device.
package media

import (
	"bufio"
	"github.com/ascherkus/go-id3/src/id3"
	"io"
	"io/ioutil"
	"os"
)

type Track struct {
	Name   string `json:name`
	Artist string `json:artist`
	Album  string `json:album`
	Path   string `json:path`
	Genre  string `json:genre`
}

// GetTracks goes trough the "../music" directory and parses all mp3's for metadata
// Returns a slice of Track structs.
func GetTracks(dir string) (*[]Track, error) {
	var Tracks []Track
	fileInfo, err := readDirectory(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range fileInfo {
		path := dir + "/" + f.Name()

		rd, err := getReader(path)
		if err != nil {
			return nil, err
		}

		id3Data := getFileInfo(rd)

		track := parseData(id3Data)
		track.Path = path

		Tracks = append(Tracks, *track)

	}

	return &Tracks, nil

}

// parseData takes i id3.File and takes out relevant data and makes a Track stucts.
// Returns a pointer to a Track struct.
func parseData(id3 *id3.File) *Track {
	track := &Track{
		Name:   id3.Name,
		Artist: id3.Artist,
		Album:  id3.Album,
		Genre:  id3.Genre,
	}

	return track
}

// readDirectory takes a string for directory and reads information about all files.
// Returns a slice of the interface os.FileInfo
func readDirectory(dir string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// getReader takes a string with a filepath to one file.
// Returns io.Reader interface.
func getReader(file string) (io.Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)
	return reader, nil

}

// getFileInfo takes io.Reader interface and finds id3 metadata from mp3s.
// Returns id3.File pointer with metadata
func getFileInfo(rd io.Reader) *id3.File {
	id3Info := id3.Read(rd)
	if id3Info == nil {
		panic("Some error")

	}
	return id3Info
}
