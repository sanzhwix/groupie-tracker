package internal

type ArtistFull struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	LocationsURL    []string            `json:"locations"`
	ConcertDatesURL []string            `json:"concertDates"`
	RelationURL     map[string][]string `json:"relations"`
}

type Event struct {
	Locations []string            `json:"locations"`
	Dates     []string            `json:"dates"`
	Relcation map[string][]string `json:"datesLocations"`
}
