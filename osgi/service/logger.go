package service

import (
	"io"
	"log"
	"strings"
)
const UNKNOWN 	= -1
const DEBUG 	= 0
const INFO 		= 1
const WARNING 	= 2
const ERROR 	= 3
// ServiceLog usable as OSGI service. The main feature is the ability to filter logs with a level. For filtering purpose, add in the message one of the key word : DEBUG, INFO, WARNING or ERROR separated by a space (' '). If none of the keyword is used, the line will not be filtered.
type ServiceLog struct {
	logger 	*log.Logger
	level 	int
	output	io.Writer
}
// Create a new ServiceLog instance with an output, a prefix, flag and the level. The three first arguments will be used to create the logger, the fourth will be used to filter the log lines.
func NewServiceLog(aOutput io.Writer, aPrefix, string, aFlag int, aLevel int) *ServiceLog {
	logger := log.New(aOutput, aPrefix, aFlags)
	service := &ServiceLog{logger : aLogger, output : aOutput, level : aLevel}
	logger.SetOutput(service) // change the output of the logger
	return service
}
// Change the log level for filter 
func (this *ServiceLog) SetLogLevel(aLogLevel int) {
	this.level = aLogLevel
}
// Obtain the log level
func (this *ServiceLog) GetLogLevel() int {
	return this.level
}
func (this *ServiceLog) GetLogger() *log.Logger {
	return this.logger
}
// Implement of writer interface
func (this *ServiceLog) Write(p []byte) (n int, err error) {
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