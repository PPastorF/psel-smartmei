package misc

import (
	"os"
	// "encoding/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ProdLogger() zerolog.Logger {
	return log.With().Caller().Logger()
}

func DevLogger() zerolog.Logger {
	return log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
}