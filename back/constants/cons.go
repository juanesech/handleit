package constants

import "os"

var DBSchema string = os.Getenv("DB_SHCEMA")
var DBAddress string = os.Getenv("DB_ADDRESS")
var DBUser string = os.Getenv("DB_USER")
var DBPass string = os.Getenv("DB_PASS")
var DBName string = os.Getenv("DB_NAME")