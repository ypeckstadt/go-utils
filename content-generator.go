package utils

import (
	"github.com/ypeckstadt/go-utils/names"
)

// Country represents country the the name needs to be generated for
type Country string

const (
	// UnitedStates generates names for the US
	UnitedStates Country = "us"

	// France generates names for France
	France Country = "france"
)


// GetRandomFirstName returns a random first name
func GetRandomFirstName(country Country) string {
	switch country {
	case UnitedStates:
		return GetRandomElement(names.UnitedStatesFirstNames)
	case France:
		return GetRandomElement(names.FranceFirstNames)
	}

	return ""
}

// GetRandomLastName returns a random last name
func GetRandomLastName(country Country) string {
	switch country {
	case UnitedStates:
		return GetRandomElement(names.UnitedStatesLastNames)
	case France:
		return GetRandomElement(names.FranceLastNames)
	}

	return ""
}

// GetRandomName returns a random first and last name
func GetRandomName(country Country) (string, string) {
	switch country {
	case UnitedStates:
		return GetRandomElement(names.UnitedStatesFirstNames), GetRandomElement(names.UnitedStatesLastNames)
	case France:
		return GetRandomElement(names.FranceFirstNames), GetRandomElement(names.FranceLastNames)
	}
	return "", ""
}


