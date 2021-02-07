package misc

import (
	"os"
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ProdLogger() zerolog.Logger {
	return log.With().Caller().Logger()
}

func DevLogger() zerolog.Logger {
	return log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func LogJsonData(msg string, d interface{}) error {
	rawData, err := json.Marshal(d);
	if err != nil {
		return err
	}

	log.Info().RawJSON("data", rawData).Msg(msg)
	return nil
}