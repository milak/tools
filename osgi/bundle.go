package osgi

import (
	"plugin"
	"github.com/google/uuid"
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
	bundleContext	*BundleContext
}
func NewPluginBundle(aPlugin *plugin.Plugin, aName string, aContext *BundleContext) Bundle {
	result := pluginBundle {id : uuid.New().String(), _plugin : aPlugin, state : RESOLVED}
	sym, err := aPlugin.Lookup("Version")
	if err == nil {
		result.version = *sym.(*string)
	} else {
		result.version = "?.?.?"
	}
	sym, err = aPlugin.Lookup("SymbolicName")
	if err == nil {
		result.symbolicName = *sym.(*string)
	} else {
		result.symbolicName = aName
	}
	return &result
}
func (this *pluginBundle) GetBundleContext() *BundleContext {
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
	if this.state == ACTIVE || this.state == STARTING{
		return
	}
	this.state = STARTING
	this.state = ACTIVE
}
func (this *pluginBundle) Stop() {
	if this.state != ACTIVE {
		return
	}
	this.state = STOPPING
	this.state = RESOLVED
}