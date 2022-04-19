package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"os"
)

type AlbumTracks []AlbumTrack
type AlbumTrack struct {
	ID    string `json:"id"`
	Items string `json: "items"`
}

func main() {
	tr, err := GetTracks("16pubXUlqRziVWRsT6lLNz")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Las canciones son ", tr)
}

func GetTracks(albumId string) (AlbumTracks, error) {
	response, err := http.Get(fmt.Sprintf("https://api.spotify.com/v1/albums/%s/tracks", albumId))
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	var tr AlbumTracks
	json.Unmarshal(bytes, &tr)
	return tr, nil
}
