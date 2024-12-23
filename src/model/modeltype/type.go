package modeltype

import "database/sql"

type Price int64
type PriceNull sql.Null[Price]

type Total int64

func GetPrice(p interface{}) int64 {
	if p == nil {
		return 0
	} else if pr, ok := p.(Price); ok {
		return int64(pr)
	} else if prn, ok := p.(PriceNull); ok {
		if prn.Valid {
			return int64(prn.V)
		}
		return 0
	} else {
		panic("not a price")
	}
}

func GetTotal(t Total) int64 {
	return int64(t)
}
