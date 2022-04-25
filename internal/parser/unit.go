package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
)

func (*fruitParser) ParseUnit(unit string) basepb.Unit {
	switch unit {
	case "KG":
		return basepb.Unit_UNIT_KG
	case "LB":
		return basepb.Unit_UNIT_LB
	default:
		log.Printf("REPO-WARNING: parser unit: %q undefined", unit)
		return basepb.Unit_UNIT_UNDEFINED
	}
}
