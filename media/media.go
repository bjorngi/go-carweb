// Package media provides support for listing and identifying music on external device.
package media

import (
	"bufio"
	vorbis "code.google.com/p/goflac-meta"
	"errors"
	"fmt"
	"github.com/ascherkus/go-id3/src/id3"
	"io"
	"io/ioutil"
	"os"
)

type Track struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album,omitempty"`
	Path   string `json:"path,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Type   string `json:"type"`
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

		switch f.Name()[len(f.Name())-3:] {
		case "mp3":
			id3Data, err := getMP3Info(path)
			if err != nil {
				return nil, err
			}

			track := parseID3(id3Data)
			track.Path = path

			Tracks = append(Tracks, *track)
		case "ogg":
			getVorbisInfo(path)

		}

	}

	return &Tracks, nil

}

// parseID3 takes id3.File and pulls out relevant data and makes a Track stucts.
// Returns a pointer to a Track struct.
func parseID3(id3 *id3.File) *Track {
	track := &Track{
		Name:   id3.Name,
		Artist: id3.Artist,
		Album:  id3.Album,
		Genre:  id3.Genre,
		Type:   "audio/mpeg",
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

func getVorbisInfo(file string) (*vorbis.Metadata, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	metadata := new(vorbis.Metadata)
	metadata.Read(f)
	fmt.Printf("%v\n", metadata)

	return metadata, nil

}

// getMP3Info takes io.Reader interface and finds id3 metadata from mp3s.
// Returns id3.File pointer with metadata
func getMP3Info(file string) (*id3.File, error) {
	rd, err := getReader(file)
	if err != nil {
		return nil, err
	}

	id3Info := id3.Read(rd)
	if id3Info == nil {
		return nil, errors.New("Missing metadata")
	}
	return id3Info, nil
}
