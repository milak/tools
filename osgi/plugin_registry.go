package osgi
import (
	"os"
	
	"log"
	"plugin"
	"github.com/milak/tools/data"
)

type Context struct {
	Logger *log.Logger
	Properties data.PropertyList
}
type pluginRegistry struct {
	pluginFolder string
	plugins []*plugin.Plugin
	context *Context
}
func NewPluginRegistry(aPluginFolder string, aContext *Context) *pluginRegistry {
	result := &pluginRegistry{pluginFolder : aPluginFolder, context : aContext}
	result.loadPlugins()
	return result
}
func (this *pluginRegistry) GetPlugins() []*plugin.Plugin {
	return this.plugins
}
func (this *pluginRegistry) GetContext() *Context {
	return this.context
}
func (this *pluginRegistry) loadPlugins(){
	// Browse plugin directory
	pluginDirectory,err := os.Open(this.pluginFolder)
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
	files,err := pluginDirectory.Readdir(0)
	if err != nil {
		this.context.Logger.Println("WARNING Unable to browse plugins directory",err)
		return
	}
	for _,file := range files {
		this.context.Logger.Println("DEBUG Loading plugin",file.Name(),"...")
		thePlugin, err := plugin.Open("plugins/"+file.Name())
		if err != nil {
			this.context.Logger.Println("WARNING Plugin has no Init method",file.Name())
		} else {
			function,err := thePlugin.Lookup("Init")
			if err != nil {
				this.context.Logger.Println("WARNING Unable to initialize plugin",file.Name(),":",err)
			} else {
				function.(func(*Context))(this.context)
			}
		}
	}
}