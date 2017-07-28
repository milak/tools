package service

import (
	"io"
	"log"
	"github.com/milak/tools/logutil"
)

// ServiceLog usable as OSGI service. The logger will use FilterableWriter from tools/logutil package 
type ServiceLog struct {
	logger 				*log.Logger
	filterableWriter	logutil.FilterableWriter
	output				io.Writer
}
// Create a new ServiceLog instance with an output, a prefix, flag and the level. The three first arguments will be used to create the logger, the fourth will be used to filter the log lines.
func NewServiceLog(aOutput io.Writer, aPrefix string, aFlag int, aLevel int) *ServiceLog {
	logger := log.New(aOutput, aPrefix, aFlag)
	service := &ServiceLog{logger : logger, output : aOutput}
	service.filterableWriter.SetLevel(aLevel)
	logger.SetOutput(this.filterableWriter) // change the output of the logger
	return service
}
// Change the log level for filter 
func (this *ServiceLog) SetLevel(aLogLevel int) {
	this.filterableWriter.SetLevel(aLogLevel)
}
// Obtain the log level
func (this *ServiceLog) GetLevel() int {
	return this.filterableWriter.GetLevel()
}
func (this *ServiceLog) GetLogger() *log.Logger {
	return this.logger
}