package parser

import (
	"github.com/marcos-wz/capstone/proto/filterpb"
	"log"
	"strings"
)

func (*fruitParser) ParseFilter(filter string) filterpb.FiltersAllowed {
	switch strings.ToUpper(filter) {
	case "ID":
		return filterpb.FiltersAllowed_FILTER_ID
	case "NAME":
		return filterpb.FiltersAllowed_FILTER_NAME
	case "COLOR":
		return filterpb.FiltersAllowed_FILTER_COLOR
	case "COUNTRY":
		return filterpb.FiltersAllowed_FILTER_COUNTRY
	default:
		log.Printf("PARSER-WARNING: filter parser : filter %q undefined", filter)
		return filterpb.FiltersAllowed_FILTER_UNDEFINED
	}
}
