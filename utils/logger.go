package utils

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
)

type logger struct {
	user     string
	filepath string
}

// Logger is Public for Exporting
var Logger logger

func (logstruct *logger) Info(logmsg string) {
	fmt.Printf("INFO: %v | USER: %v", logmsg, logstruct.user)
}

func (logstruct *logger) Warning(logmsg string) {
	fmt.Printf("WARNING: %v | USER: %v", logmsg, logstruct.user)
}

func (logstruct *logger) Error(logmsg string) {
	fmt.Printf("ERROR: %v | USER: %v", logmsg, logstruct.user)
}

func init() {
	// Set Up Logger Attr's

	user, err := user.Current()
	if err != nil {
		log.Fatal("Cannot recognize user")
	}
	Logger.user = user.Username
	abspath, err := filepath.Abs("../logs") //Run our of src directory for src/logs
	if err != nil {
		log.Fatal("Error Finding logs Directory")
	}
	logpath := abspath + "/earningsreport-log.log"
	Logger.filepath = logpath
}
