package service

import (
	"io"
	"log"
	"os"
)
const UNKNOWN 	= -1
const DEBUG 	= 0
const INFO 		= 1
const WARNING 	= 2
const ERROR 	= 3
type ServiceLog struct {
	log.Logger
	level 	int
	out		io.Writer
}
func NewServiceLog(prefix string,flag int, aOutput io.Writer){
	logger := log.New(os.Stdout, prefix, flag)
	this.out = aOutput
	logger.SetOutput(this)
}
func (this *ServiceLog) SetLogLevel(aLogLevel int){
	this.level = aLogLevel
}
func (this *ServiceLog) GetLogLevel(){
	return this.level
}
// Implement of writer interface
func (this *ServiceLog) Write(p []byte) (n int, err error) {
	levelName := ""
	i := 0
	for (i < len(p) && p[i] != ' ') {
		i++
		levelName += p[i]
	}
	var level int
	if i == len(p) {
		level = levelFromName(levelName)
	} else {
		level = UNKNOWN
	}
	if level == UNKNOWN {
		this.out.Write(p)
	if this.level >= level {
		this.out.Write(p[i:])
	} else {
		// Filtred
	}
}
func levelFromName(levelName) int {
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