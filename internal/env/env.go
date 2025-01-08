package env

import (
	"os"
	"strconv"
)

var Env = struct {
	Port            int
	AppEnvironment  string
	DbConnectionStr string
}{
	Port:            getIntVal("PORT", 5000),
	AppEnvironment:  getStringVal("ENV", "development"),
	DbConnectionStr: getStringVal("DB_CONN_STR", "postgresql://msvc_user:msvc_password@localhost:5431/msvc_db?sslmode=disable"),
}

func getStringVal(prop string, fallback string) string {
	if val, ok := os.LookupEnv(prop); ok {
		return val
	}

	return fallback
}

func getIntVal(prop string, fallback int) int {
	if val, ok := os.LookupEnv(prop); ok {
		if str, err := strconv.Atoi(val); err == nil {
			return str
		}
	}

	return fallback
}
