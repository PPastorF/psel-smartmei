package misc

import (
	"time"
	"strings"
	"github.com/google/uuid"
)

func GenerateUniqueID() UniqueID {
	return UniqueID(uuid.New().String())
}

func TimestampStr(t *time.Time) (string, error) {
	ts, err := time.Now().MarshalText()
	if err != nil {
		return "", err
	}

	return string(ts), nil
}

func SanitizeString(raw string) (string, error) {
	s := strings.TrimSpace(raw)	
	
	return s, nil
}