package media

import (
	"fmt"
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

func TestGetMP3Info(t *testing.T) {
	testfile := "../music/06 Hasta Siempre.mp3"
	//f, err := os.Open(testfile)
	//if err != nil {
	//	t.Errorf("Failed to open: %v\n", testfile)
	//}

	id3Info, err := getMP3Info(testfile)
	if err != nil {
		t.Errorf("Failed getting MP3Info")
	}

	if id3Info.Artist != "Antamerica" {
		t.Errorf("Wrong metadata, was expeting \"Hasta Siempre\" got \"%v\"", id3Info.Artist)

	}

	if id3Info.Name != "Hasta Siempre" {
		t.Errorf("Wrong metadata, was expeting \"Human\" got \"%v\"", id3Info.Name)

	}

	fmt.Printf("%+v\n\n", id3Info)

}

func TestGetTracks(t *testing.T) {
	tracks, err := GetTracks("../music")
	if err != nil {
		t.Errorf("Failed getting all tracks: ", err)
	}

	fmt.Printf("TestGetTracks: %+v\n\n", tracks)
}

func TestGetVorbisInfo(t *testing.T) {
	testfile := "../music/test.ogg"
	metadata, err := getVorbisInfo(testfile)
	if err != nil {
		t.Errorf("Failed to read vorbis metadata: %v\n", err)
	}
	fmt.Printf("%+v\n", metadata)
}
