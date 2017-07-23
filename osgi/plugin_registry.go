package osgi

import (
	"os"
	"github.com/milak/tools/data"
	"log"
	"plugin"
)
// The OSGI context class
type Context struct {
	Logger     *log.Logger
	Properties data.PropertyList
}
// The PluginRegistry class
type pluginRegistry struct {
	pluginFolder string
	plugins      []*plugin.Plugin
	context      *Context
}
// Create a NewPluginRegistry with a folder name containing the plugins and an initialized context. 
// Once created, the registry will load the plugins.
func NewPluginRegistry(aPluginFolder string, aContext *Context) *pluginRegistry {
	result := &pluginRegistry{pluginFolder: aPluginFolder, context: aContext}
	result.loadPlugins()
	return result
}
// Obtain the list of the loaded plugins
func (this *pluginRegistry) GetPlugins() []*plugin.Plugin {
	return this.plugins
}
// Obtain the context
func (this *pluginRegistry) GetContext() *Context {
	return this.context
}
// Load the plugins in the plugin folder
func (this *pluginRegistry) loadPlugins() {
	// Browse plugin directory
	pluginDirectory, err := os.Open(this.pluginFolder)
	if err != nil {
		this.context.Logger.Println("WARNING No plugin directory")
		// no plugins directory
		return
	}
	this.context.Logger.Println("INFO Loading plugins...")
	info, err := pluginDirectory.Stat()
	if !info.IsDir() {
		this.context.Logger.Println("WARNING Plugins directory is not a directory")
		return
	}
	files, err := pluginDirectory.Readdir(0)
	if err != nil {
		this.context.Logger.Println("WARNING Unable to browse plugins directory", err)
		return
	}
	for _, file := range files {
		this.loadPlugin(file)
	}
}
// Load one plugin
func (this *pluginRegistry) loadPlugin(file os.FileInfo) {
	defer func() {
		if r := recover(); r != nil {
			this.context.Logger.Println("WARNING Failed to initialize plugin", file.Name(), ":", r)
		}
	}()
	this.context.Logger.Println("DEBUG Loading plugin", file.Name(), "...")
	thePlugin, err := plugin.Open("plugins/" + file.Name())
	if err != nil {
		this.context.Logger.Println("WARNING Plugin has no Init method", file.Name())
	} else {
		this.plugins = append(this.plugins,thePlugin)
		function, err := thePlugin.Lookup("Init")
		if err != nil {
			this.context.Logger.Println("WARNING Unable to initialize plugin", file.Name(), ":", err)
		} else {
			function.(func(*Context))(this.context)
		}
	}
}
