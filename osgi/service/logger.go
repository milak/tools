package service

import (
	"io"
	"log"
	"os"
	"fmt"
)
const UNKNOWN 	= -1
const DEBUG 	= 0
const INFO 		= 1
const WARNING 	= 2
const ERROR 	= 3
type ServiceLog struct {
	logger 	*log.Logger
	level 	int
	output	io.Writer
}
func NewServiceLog(prefix string,flag int, aOutput io.Writer) *ServiceLog {
	logger := log.New(os.Stdout, prefix, flag)
	service := &ServiceLog{logger : logger, output : aOutput, level : INFO}
	logger.SetOutput(service)
	return service
}
func (this *ServiceLog) SetLogLevel(aLogLevel int){
	this.level = aLogLevel
}
func (this *ServiceLog) GetLogLevel() int{
	return this.level
}
// Implement of writer interface
func (this *ServiceLog) Write(p []byte) (n int, err error) {
	levelName := ""
	i := 0
	for (i < len(p) && p[i] != ' ') {
		levelName += string(p[i])
		i++
	}
	var level int
	if i != len(p) {
		level = levelFromName(levelName)
	} else {
		level = UNKNOWN
	}
	if level == UNKNOWN {
		this.output.Write(p)
	} else if this.level >= level {
		this.output.Write(p[i:])
		this.output.Write([]byte("\n"))
	} else {
		// Filtred
	}
	return len(p),nil
}
func levelFromName(levelName string) int {
	if levelName == "DEBUG" {
		return DEBUG
	} else if levelName == "INFO" {
		return INFO
	} else if levelName == "WARNING" {
		return WARNING
	} else if levelName == "ERROR" {
		return ERROR
	} else {
		return UNKNOWN
	}
}