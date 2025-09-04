package internal

type ArtistFull struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string

	Locations []string
	Dates     []string
	Relation  map[string][]string
}
