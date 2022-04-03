package main

import (
	"errors"
	"github.com/rs/zerolog/log"
)

func main() {
	err := errors.New("there is an error")

	log.Error().Err(err).Msg("this the way to log errors")
}
