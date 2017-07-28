package osgi

import (
	"log"
	"plugin"
	"github.com/google/uuid"
	"github.com/milak/tools/osgi/service"
	"github.com/milak/tools/logutil"
)
const UNINSTALLED = 1
const INSTALLED   = 2
const RESOLVED    = 3
const STARTING    = 4
const STOPPING    = 5
const ACTIVE      = 6

type Bundle interface {
	GetBundleId() string
	GetState() int
	GetSymbolicName() string
	GetVersion() string
	Start()
	Stop()
}
type pluginBundle struct {
	id 				string
	version 		string
	symbolicName 	string
	_plugin 		*plugin.Plugin
	state 			int
	bundleContext	BundleContext
	logger			*log.Logger
}
func NewPluginBundle(aPlugin *plugin.Plugin, aName string, aContext BundleContext) Bundle {
	result := pluginBundle {id : uuid.New().String(), _plugin : aPlugin, state : RESOLVED}
	// Getting logger 
	logServiceRef := aContext.GetService("LogService")
	if logServiceRef != nil {
		logService := logServiceRef.Get().(service.LogService)
		result.logger = logService.GetLogger()
	} else {
		result.logger = logutil.DefaultLogger
		result.logger.Println("Using default logger")
	}
	
	sym, err := aPlugin.Lookup("Version")
	if err == nil {
		result.version = *sym.(*string)
	} else {
		result.version = "?.?.?"
		result.logger.Println(err)
	}
	sym, err = aPlugin.Lookup("SymbolicName")
	if err == nil {
		result.symbolicName = *sym.(*string)
	} else {
		result.symbolicName = aName
	}
	return &result
}
func (this *pluginBundle) GetBundleContext() BundleContext {
	return this.bundleContext
}
func (this *pluginBundle) GetBundleId() string {
	return this.id
}
func (this *pluginBundle) GetSymbolicName() string {
	return this.symbolicName
}
func (this *pluginBundle) GetVersion() string {
	return this.version
}
func (this *pluginBundle) GetState() int {
	return this.state
}
func (this *pluginBundle) Start() {
	if this.state != INSTALLED {
		return
	}
	this.state = STARTING
	defer func() {
		if r := recover(); r != nil {
			this.logger.Println("WARNING Failed to Start bundle", this.symbolicName, ":", r)
			this.state = INSTALLED
		}
	}()
	function, err := thePlugin.Lookup("Start")
	if err != nil {
		this.logger.Println("WARNING Unable to initialize plugin", this.symbolicName, ":", err)
	} else {
		function.(func(BundleContext))(context)
	}
	this.state = ACTIVE
}
func (this *pluginBundle) Stop() {
	if this.state != ACTIVE {
		return
	}
	this.state = STOPPING
	defer func() {
		if r := recover(); r != nil {
			this.logger.Println("WARNING Failed to Stop bundle", this.symbolicName, ":", r)
			this.state = INSTALLED
		}
	}()
	function, err := thePlugin.Lookup("Stop")
	if err != nil {
		this.logger.Println("WARNING Unable to initialize plugin", this.symbolicName, ":", err)
	} else {
		function.(func(BundleContext))(context)
	}
	this.state = RESOLVED
}