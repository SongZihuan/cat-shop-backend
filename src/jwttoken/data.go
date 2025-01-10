package jwttoken

import "time"

type Data struct {
	userid    uint
	resetmin  time.Duration
	resettime time.Time
}

func (d *Data) Userid() uint {
	return d.userid
}

func (d *Data) ResetMin() time.Duration {
	return d.resetmin
}

func (d *Data) IsNowReset() bool {
	if d.resetmin > 0 {
		return time.Now().After(d.resettime)
	}
	return false
}
