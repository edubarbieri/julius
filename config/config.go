package config

import (
	"os"
	"strconv"
)

var (
	PostgresURL string
	JwtSecret   string
	HttpPort    int
)

func init() {
	PostgresURL = os.Getenv("POSTGRES_URL")
	JwtSecret = os.Getenv("JWT_SECRET")
	HttpPort = getIntValue("PORT", 8081)
}

func getIntValue(varName string, defaultValue int) int {
	varValue := os.Getenv(varName)
	intValue, err := strconv.Atoi(varValue)
	if err != nil {
		return defaultValue
	}
	return intValue
}
