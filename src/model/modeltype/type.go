package modeltype

import "database/sql"

type Price uint64

func (p Price) ToInt64() int64 {
	return int64(p)
}

type PriceNull sql.Null[Price]

func (p PriceNull) ToPrice() Price {
	if !p.Valid {
		return Price(0)
	}
	return p.V
}

func (p PriceNull) ToInt64() int64 {
	return int64(p.ToPrice())
}

type PriceNullJson int64

func (p PriceNullJson) ToPrice() Price {
	if p <= 0 {
		return Price(0)
	}
	return Price(p)
}

func (p PriceNullJson) ToPriceNull() PriceNull {
	return PriceNull{V: p.ToPrice(), Valid: p >= 0}
}

func (p PriceNullJson) ToInt64() int64 {
	return int64(p.ToPrice())
}

type Total uint64

func (t Total) ToInt64() int64 {
	return int64(t)
}
