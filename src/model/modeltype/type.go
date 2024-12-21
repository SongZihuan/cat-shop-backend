package modeltype

type Price int64
type PriceNull *int64

type Total int64

func GetPrice(p interface{}) int64 {
	if p == nil {
		return 0
	} else if pr, ok := p.(Price); ok {
		return int64(pr)
	} else if prn, ok := p.(PriceNull); ok {
		if prn == nil {
			return 0
		}
		return *prn
	} else {
		panic("not a price")
	}
}

func GetTotal(t Total) int64 {
	return int64(t)
}
