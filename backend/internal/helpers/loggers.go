package helpers

import (
	"log"
	"os"
)

var InfoLogger = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
var ErrorLogger = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
