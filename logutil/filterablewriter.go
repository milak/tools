package logutil

import (
	"io"
)
const UNKNOWN 	= -1
const DEBUG 	= 0
const INFO 		= 1
const WARNING 	= 2
const ERROR 	= 3
// The feature of FilterableWriter is the ability to filter logs with a level. For filtering purpose, add in the message one of the key word : DEBUG, INFO, WARNING or ERROR separated by a space (' '). If none of the keyword is used, the line will not be filtered.
type FilterableWriter struct {
	level 	int
	output	io.Writer
}
func NewFilterableWriter(aLevel int, aOutput io.Writer) *FilterableWriter{
	return &FilterableWriter{level : aLevel, output : aOutput}
}
func (this *FilterableWriter) GetLevel() {
	return this.level
}
func (this *FilterableWriter) SetLevel(aLevel int) {
	this.level = aLevel
}
// Implement of writer interface
func (this *FilterableWriter) Write(p []byte) (n int, err error) {
	// TODO perfomance leaks change implementation or use thread to avoid slowing the process 
	logLine := string(p)
	var level int
	if strings.Contains(logLine, "DEBUG ") {
		level = DEBUG
	} else if strings.Contains(logLine, "INFO ") {
		level = INFO
	} else if strings.Contains(logLine, "WARNING ") {
		level = WARNING
	} else if strings.Contains(logLine, "ERROR ") {
		level = ERROR
	} else {
		level = UNKNOWN
	}
	if level == UNKNOWN {
		this.output.Write(p)
	} else if this.level <= level {
		this.output.Write(p)
	} else {
		// Filtered
	}
	return len(p),nil
}