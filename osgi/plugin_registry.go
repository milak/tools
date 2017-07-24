package osgi

import (
	"os"
	"github.com/milak/tools/data"
	"log"
	"plugin"
)

// The framework class
type framework struct {
	bundleFolder string
	bundles      []*Bundle
	context      *BundleContext
	Logger     *log.Logger
	Properties data.PropertyList
}
// Create a NewPluginRegistry with a folder name containing the plugins and an initialized context. 
// Once created, the registry will load the plugins.
func NewFramework(aBundleFolder string, aContext *BundleContext) *pluginRegistry {
	result := &framework{bundleFolder: aBundleFolder, context: aContext}
	result.loadBundles()
	return result
}
// Obtain the list of the loaded plugins
func (this *framework) GetBundles() []*Bundle {
	return this.bundles
}
// Obtain the context
func (this *framework) GetContext() *BundleContext {
	return this.context
}
// Load the plugins in the plugin folder
func (this *framework) loadBundles() {
	// Browse plugin directory
	bundleDirectory, err := os.Open(this.bundleFolder)
	if err != nil {
		this.context.Logger.Println("WARNING Bundle directory doesn't exist")
		// no bundle directory
		return
	}
	this.context.Logger.Println("INFO Loading bundles...")
	info, err := bundleDirectory.Stat()
	if !info.IsDir() {
		this.context.Logger.Println("WARNING Bundles directory is not a directory")
		return
	}
	files, err := bundleDirectory.Readdir(0)
	if err != nil {
		this.context.Logger.Println("WARNING Unable to browse bundle directory", err)
		return
	}
	for _, file := range files {
		this.loadBundle(file)
	}
}
// Load one plugin
func (this *framework) loadBundle(file os.FileInfo) {
	defer func() {
		if r := recover(); r != nil {
			this.context.Logger.Println("WARNING Failed to initialize bundle", file.Name(), ":", r)
		}
	}()
	this.context.Logger.Println("DEBUG Loading bundle", file.Name(), "...")
	thePlugin, err := plugin.Open("plugins/" + file.Name())
	if err != nil {
		this.context.Logger.Println("WARNING Bundle has no Init method", file.Name())
	} else {
		bundle := NewPluginBundle(thePlugin, file.Name(), this.context)
		this.plugins = append(this.bundles,bundle)
		function, err := thePlugin.Lookup("Init")
		if err != nil {
			this.context.Logger.Println("WARNING Unable to initialize plugin", file.Name(), ":", err)
		} else {
			function.(func(*Context))(this.context)
		}
	}
}