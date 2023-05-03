package utils

import (
	"encoding/json"

	"github.com/rs/zerolog"
)

func ResponseBuilder(output any, logger *zerolog.Logger) string {
	res, err := json.Marshal(output)
	if err != nil {
		logger.Info().Msg(err.Error())
		return ""
	}

	logger.Info().Msg(string(res))
	return string(res)
}
