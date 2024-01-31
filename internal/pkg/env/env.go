package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func GetEnvAsBool(name string, defaultVal bool) bool {
	valueStr := GetEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valueStr := GetEnv(name, "")

	if valueStr == "" {
		return defaultVal
	}

	return strings.Split(valueStr, sep)
}

func GetEnvTimeDuration(name string, defaultVal int) time.Duration {
	value := GetEnvAsInt(name, defaultVal)
	return time.Duration(value)
}
