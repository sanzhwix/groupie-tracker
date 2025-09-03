package application

import (
	"groupie-tracker/internal/"
)

type ArtistService struct {}


func (s *ArtistService) GetAllArtist() ([]domian.Artist, error) {
	return api.FetchArtist()
}