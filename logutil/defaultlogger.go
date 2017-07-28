package logutil

import (
	"log"
	"os"
)

var DefaultLogger = log.New(os.Stderr, "", log.LstdFlags)