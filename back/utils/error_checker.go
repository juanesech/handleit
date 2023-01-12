package utils

import (
	"github.com/juanesech/topo/constants"
	log "github.com/sirupsen/logrus"
)

// Log on error
func CheckError(err error) {
	if err != nil {
		if err.Error() == constants.DBName {
			log.Error("Database not found")
		}
		log.Error(err)
	}
}
