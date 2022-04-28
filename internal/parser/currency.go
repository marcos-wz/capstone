package parser

import (
	"github.com/marcos-wz/capstone/proto/basepb"
	"log"
)

func (*fruitParser) ParseCurrency(currency string) basepb.Currency {
	switch currency {
	case "MXN":
		return basepb.Currency_CURRENCY_MXN
	case "BRL":
		return basepb.Currency_CURRENCY_BRL
	case "CAD":
		return basepb.Currency_CURRENCY_CAD
	case "USD":
		return basepb.Currency_CURRENCY_USD
	default:
		log.Printf("PARSER-WARNING: currency parser : currency %q undefined", currency)
		return basepb.Currency_CURRENCY_UNDEFINED
	}
}
