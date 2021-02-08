package misc

import (
	"os"
	"fmt"
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

func InProductionEnv() bool {
    if strings.ToUpper(os.Getenv("DEPLOY_ENV")) == "PROD" {
        return true
    } else {
        return false
    }
}

func ApiRoutePrefix(version string) string {
	return fmt.Sprintf("/api/%s", version)
}

func ConnectionURL(addr, port string) string {
	return fmt.Sprintf("%s:%s", addr, port)
}

func SanitizeString(raw string) string {
	s := strings.TrimSpace(raw)	
	
	return s
}