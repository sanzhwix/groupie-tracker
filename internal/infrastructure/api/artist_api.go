package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker/internal/domain"
)


const artistAPI = "https://groupietrackers.herokuapp.com/api/artists"


func FetchArtist() ([]domain.Artist, error) {
	resp, err := http.Get(artistAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status is not OK: %d", resp.StatusCode)
	}

	var artists []domain.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, fmt.Errorf("Failed to deocde JSON file", err)
	}

	return artist, nil
}






/*
1) Read the api body
2) close
3) check the status code 
4) read json and fill it 
\*