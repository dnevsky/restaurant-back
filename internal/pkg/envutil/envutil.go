package envutil

import (
	"os"
	"strconv"
	"time"
)

// GetString возвращает значение переменной окружения по ключу или значение по умолчанию, если переменная не установлена.
func GetString(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// GetInt возвращает значение переменной окружения, преобразованное в int, или значение по умолчанию, если преобразование не удалось.
func GetInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if intValue, err := strconv.Atoi(val); err == nil {
			return intValue
		}
	}
	return defaultVal
}

// GetFloat возвращает значение переменной окружения, преобразованное в float64, или значение по умолчанию, если преобразование не удалось.
func GetFloat(key string, defaultVal float64) float64 {
	if val := os.Getenv(key); val != "" {
		if floatValue, err := strconv.ParseFloat(val, 64); err == nil {
			return floatValue
		}
	}
	return defaultVal
}

// GetBool возвращает значение переменной окружения, преобразованное в bool, или значение по умолчанию, если преобразование не удалось.
func GetBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if boolValue, err := strconv.ParseBool(val); err == nil {
			return boolValue
		}
	}
	return defaultVal
}

// GetDuration возвращает значение переменной окружения, преобразованное в time.Duration,
// или значение по умолчанию, если переменная не установлена или преобразование не удалось.
// Формат строки должен соответствовать формату, поддерживаемому функцией time.ParseDuration (например, "300ms", "1.5h", "2h45m").
func GetDuration(key string, defaultVal time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			return duration
		}
	}
	return defaultVal
}
