package service

import (
	"io"
	"log"
	"github.com/milak/tools/logutil"
	"os"
)

// LogService usable as OSGI service. The logger will use FilterableWriter from tools/logutil package 
type LogService struct {
	logger 				*log.Logger
	filterableWriter	*logutil.FilterableWriter
}
func NewDefaultLogService() *LogService {
	return NewLogService(os.Stdout, "",  log.Ldate | log.Ltime | log.Lshortfile, logutil.INFO)
}
// Create a new LogService instance with an output, a prefix, flag and the level. The three first arguments will be used to create the logger, the fourth will be used to filter the log lines.
func NewLogService(aOutput io.Writer, aPrefix string, aFlag int, aLevel int) *LogService {
	logger := log.New(aOutput, aPrefix, aFlag)
	service := &LogService{logger : logger}
	service.filterableWriter = logutil.NewFilterableWriter(aLevel,aOutput)
	logger.SetOutput(service.filterableWriter) // change the output of the logger
	return service
}
// Change the log level for filter 
func (this *LogService) SetLevel(aLogLevel int) {
	this.filterableWriter.SetLevel(aLogLevel)
}
// Obtain the log level
func (this *LogService) GetLevel() int {
	return this.filterableWriter.GetLevel()
}
func (this *LogService) GetLogger() *log.Logger {
	return this.logger
}