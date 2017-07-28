package logutil

import (
	"os"
	"log"
)

func ExampleFilterableWriter() {
	writer := NewFilterableWriter(WARNING,os.Stdout)
	log.SetOutput(writer)
	log.Println("DEBUG hello") // will be filtered
	log.Println("WARNING hello") // will not be filtered
	log.Println("hello") // will not be filtered
	writer.SetLevel(DEBUG)
	log.Println("DEBUG hello") // will not be filtered
}