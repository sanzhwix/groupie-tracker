package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ArtistApi = "https://groupietrackers.herokuapp.com/api/artists"

func FetchArtist() ([]ArtistFull, error) {
	resp, err := http.Get(ArtistApi)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not the OK! %d", resp.StatusCode)
	}

	var artistsInfo []ArtistFull
	err = json.NewDecoder(resp.Body).Decode(&artistsInfo)
	if err != nil {
		return nil, fmt.Errorf("error with decoding JSON %v", err)
	}

	return artistsInfo, nil
}

// LOCATION FETCHING
const LocationsApi = "https://groupietrackers.herokuapp.com/api/locations"

type LocationsResponse struct {
	ID            int      `json:"id"`
	LocationsResp []string `json:"locations"`
	// DatesURL      string   `json:"dates"`
}

func FetchLocactions() ([]string, error) {
	resp, err := http.Get(LocationsApi)
	if err != nil {
		return nil, fmt.Errorf("error fetching locations! %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status is not okay in locations %v", resp.StatusCode)
	}

	var locResp LocationsResponse
	err = json.NewDecoder(resp.Body).Decode(&locResp)
	if err != nil {
		return nil, fmt.Errorf("error decoding locations! %v", err)
	}

	return locResp.LocationsResp, nil
}

// FETCING DATES
const DatesApi = "https://groupietrackers.herokuapp.com/api/dates"

type DatesResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func FetchDates() ([]string, error) {
	resp, err := http.Get(DatesApi)
	if err != nil {
		return nil, fmt.Errorf("error caused fethcing Dates! %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response code is not ok in dates! %v", resp.StatusCode)
	}

	var datesResp DatesResponse
	err = json.NewDecoder(resp.Body).Decode(&datesResp)
	if err != nil {
		return nil, fmt.Errorf("somethinf went wrong while decoding dates! %v", err)
	}

	return datesResp.Dates, nil
}
