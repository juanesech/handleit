package utils

import (
	log "github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		log.Error(err)
	}
}
