package utils

import (
	"database/sql"
	"regexp"
)

const BASE_CHAR = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStr(length int) string {
	bytes := []byte(BASE_CHAR)

	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, bytes[Rand().Intn(len(bytes))])
	}

	return string(result)
}

func IsChinaMainlandPhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

func GetSQLNullString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
