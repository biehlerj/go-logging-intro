package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Debug().Msg("Debug message is displayed")
	log.Info().Msg("Info message is displayed")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msg("Debug message is no longer displayed")
	log.Info().Msg("Info message is displayed")
}
