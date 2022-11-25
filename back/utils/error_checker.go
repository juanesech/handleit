package utils

import (
	"github.com/juanesech/topo/constants"
	log "github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		if err.Error() == constants.DBName {
			log.Error("Database not found")
		}
		log.Error(err)
	}
}
