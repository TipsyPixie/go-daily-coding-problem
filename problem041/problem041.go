package problem041

type airport struct {
	next *airport
	code string
}

func (currentAirport *airport) tryAddNext(nextAirport *airport) bool {
	if currentAirport.next == nil || currentAirport.next.code > nextAirport.code {
		currentAirport.next = nextAirport
		return true
	}
	return false
}

func FullItinerary(flights [][]string, startingAirportCode string) []string {
	airportMap := map[string]*airport{}
	startingAirportMap := map[string]*airport{}

	for _, flight := range flights {
		originAirportCode, destinationAirportCode := flight[0], flight[1]

		if airportMap[destinationAirportCode] == nil {
			newDestinationAirport := airport{
				next: nil,
				code: destinationAirportCode,
			}
			airportMap[destinationAirportCode] = &newDestinationAirport
		} else {
			startingAirportMap[destinationAirportCode] = nil
		}

		if airportMap[originAirportCode] == nil {
			newOriginAirport := airport{
				next: airportMap[destinationAirportCode],
				code: originAirportCode,
			}
			startingAirportMap[originAirportCode], airportMap[originAirportCode] = &newOriginAirport, &newOriginAirport
		} else {
			airportMap[originAirportCode].tryAddNext(airportMap[destinationAirportCode])
		}
	}

	startingAirport := startingAirportMap[startingAirportCode]
	if startingAirport == nil {
		return nil
	}
	itinerary := make([]string, 0, 4)
	for currentAirport := startingAirport; currentAirport != nil; currentAirport = currentAirport.next {
		itinerary = append(itinerary, currentAirport.code)
	}
	return itinerary
}
