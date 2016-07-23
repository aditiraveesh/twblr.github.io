package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	superheros_with_interest := []string{}

	for superhero, interests := range data {
		if contains(interests, interest) {
			superheros_with_interest = append(superheros_with_interest, superhero)	
		}
	}

	return superheros_with_interest
}

func contains(src []string, elem string) bool {
	for _, str := range src {
		if str == elem {
			return true
		}
	}
	return false
}
