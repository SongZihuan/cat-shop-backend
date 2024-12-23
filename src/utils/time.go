package utils

import (
	"database/sql"
	"time"
)

func GetSQLNullTimeUnix(s sql.NullTime) int64 {
	if s.Valid {
		return s.Time.Unix()
	}
	return 0
}

func SqlNullNow() sql.NullTime {
	return sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
}
