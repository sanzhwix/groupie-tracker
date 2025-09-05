package internal

func EventBuilder(a Artist) (Event, error) {
	locs, err := FetchLocactions(a.LocationsURL)
	if err != nil {
		return Event{}, err
	}

	dates, err := FetchDates(a.ConcertDatesURL)
	if err != nil {
		return Event{}, err
	}

	rels, err := FetchRelations(a.RelationURL)
	if err != nil {
		return Event{}, err
	}

	return Event{
		Locations: locs,
		Dates:     dates,
		Relcation: rels,
	}, nil
}
