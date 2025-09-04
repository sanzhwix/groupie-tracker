package main

import (
	"fmt"
	"groupie-tracker/internal"
)

func main() {
	artists, err := internal.FetchArtist()
	if err != nil {
		fmt.Println("Error fetching ArtistsFull!")
		return
	}

	for i := 0; i <= 5; i++ {
		fmt.Printf("%d: %s (%d)\n", artists[i].ID, artists[i].Name, artists[i].CreationDate)
	}
}
