package conv

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

// Для конвертации QuetyParam из строки в число
func ConvertQuery(query string) int {
	num, err := strconv.Atoi(query)
	if err != nil {
		logrus.Errorf("Err convert %s to type int: %s", query, err)
		return 0
	}
	return num
}
