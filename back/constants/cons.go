package constants

import "os"

var DBAddress string = os.Getenv("DB_ADDRESS")
var DBName string = os.Getenv("DB_NAME")
