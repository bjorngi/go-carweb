package media

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// Test reading directory
func TestReadDir(t *testing.T) {
	files, err := readDirectory("../music")
	if err != nil {
		t.Errorf("Reading directory failed: %v\n", err)
	}

	fmt.Printf("%d songs in directory\n\n", len(files))
}

func TestGetFileInfo(t *testing.T) {
	testfile := "../music/The Killers-Human.mp3"
	f, err := os.Open(testfile)
	if err != nil {
		t.Errorf("Failed to open: %v\n", testfile)
	}

	reader := bufio.NewReader(f)

	id3Info := getFileInfo(reader)

	if id3Info.Artist != "The Killers" {
		t.Errorf("Wrong metadata, was expeting \"The Killers\" got \"%v\"", id3Info.Name)

	}

	if id3Info.Name != "Human" {
		t.Errorf("Wrong metadata, was expeting \"Human\" got \"%v\"", id3Info.Name)

	}

	fmt.Printf("%+v\n\n", id3Info)

}

func TestGetTracks(t *testing.T) {
	tracks, err := GetTracks("../music")
	if err != nil {
		t.Errorf("Failed getting all tracks: ", err)
	}

	fmt.Printf("%+v", tracks)
}
