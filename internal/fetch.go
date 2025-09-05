package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ArtistApi = "https://groupietrackers.herokuapp.com/api/artists"

func FetchArtist() ([]Artist, error) {
	resp, err := http.Get(ArtistApi)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not the OK! %d", resp.StatusCode)
	}

	var artistsInfo []Artist
	err = json.NewDecoder(resp.Body).Decode(&artistsInfo)
	if err != nil {
		return nil, fmt.Errorf("error with decoding JSON %v", err)
	}

	return artistsInfo, nil
}

// LOCATION FETCHING
// const LocationsApi = "https://groupietrackers.herokuapp.com/api/locations"

type LocationsResponse struct {
	ID            int      `json:"id"`
	LocationsResp []string `json:"locations"`
	// DatesURL      string   `json:"dates"`
}

func FetchLocactions(url string) ([]string, error) {
	resp, err := http.Get(url)
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
// const DatesApi = "https://groupietrackers.herokuapp.com/api/dates"

type DatesResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func FetchDates(url string) ([]string, error) {
	resp, err := http.Get(url)
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

// FETCHING RELATIONS

// const RelationsApi = "https://groupietrackers.herokuapp.com/api/relation"

type RelationsResp struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchRelations(url string) (map[string][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error with fethcing relations %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status is not ok in relations %v", resp.StatusCode)
	}

	var Rel RelationsResp
	err = json.NewDecoder(resp.Body).Decode(&Rel)
	if err != nil {
		return nil, fmt.Errorf("erorr fetching relations! %v", err)
	}

	return Rel.DatesLocations, nil
}
