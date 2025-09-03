package application

import (
	"groupie-tracker/internal/domain"
	"groupie-tracker/internal/infrastructure/api"
)

type ArtistService struct {}


func (s *ArtistService) GetAllArtist() ([]domian.Artist, error) {
	return api.FetchArtist()
}