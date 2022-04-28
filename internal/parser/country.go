package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
	"strings"
)

func (*fruitParser) ParseCountry(country string) basepb.Country {
	switch strings.ToUpper(country) {
	case "MEXICO":
		return basepb.Country_COUNTRY_MEXICO
	case "BRAZIL":
		return basepb.Country_COUNTRY_BRAZIL
	case "CANADA":
		return basepb.Country_COUNTRY_CANADA
	case "USA":
		return basepb.Country_COUNTRY_USA
	default:
		log.Printf("PARSER-WARNING: country parser: country %q undefined", country)
		return basepb.Country_COUNTRY_UNDEFINED
	}
}
