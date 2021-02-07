package misc

import (
	"time"
    types "github.com/ppastorf/psel-smartmei/internal/types"
)

func GenerateTimestamp() (types.Timestamp, error) {
	ts, err := time.Now().MarshalText()
	if err != nil {
		return "", err
	}

	return types.Timestamp(ts), nil
}