package osgi

import (
	"os"
	"github.com/milak/tools/osgi/service"
	"github.com/milak/tools/data"
	"github.com/google/uuid"
	"log"
	"plugin"
)

// The framework class
type Framework struct {
	bundleId		string
	bundleFolder 	string
	bundles      	[]Bundle
	logger     		*log.Logger
	properties 		data.PropertyList
	services		map[string]*ServiceRef
	state			int
}
// Create a NewPluginRegistry with a folder name containing the plugins and an initialized context. 
// Once created, the registry will load the plugins.
func NewFramework(aBundleFolder string) *Framework {
	result := &Framework{bundleFolder: aBundleFolder}
	result.bundleId = uuid.New().String()
	result.services = make(map[string]*ServiceRef)
	return result
}
func (this *Framework) GetBundleId() string {
	return this.bundleId
}
func (this *Framework) GetSymbolicName() string {
	return "framework"
}
func (this *Framework) GetVersion() string {
	return "1.0"
}
func (this *Framework) Start(){
	if this.state == ACTIVE || this.state == STARTING{
		return
	}
	// Looking for LogService
	logServiceRef := this.GetService("LogService")
	if logServiceRef == nil {
		logService := service.NewDefaultServiceLog()
		this.RegisterService("LogService", &logService)
		this.logger = logService.GetLogger()
	} else {
		logService := logServiceRef.Get().(*service.LogService)
		this.logger = logService.GetLogger()
	}
	// Starting
	this.logger.Println("INFO Starting...")
	this.state = STARTING
	this.loadBundles()
	this.state = ACTIVE
	this.logger.Println("INFO Active")
}
func (this *Framework) Stop(){
	if this.state != ACTIVE {
		return
	}
	this.logger.Println("INFO Stopping...")
	this.state = STOPPING
	for _,bundle := range this.bundles {
		bundle.Stop()
	}
	this.state = RESOLVED
	this.logger.Println("INFO Resolved")
}
func (this *Framework) GetState() int {
	return this.state
}
func (this *Framework) GetProperty(aName string) interface{} {
	if !this.properties.HasProperty(aName) {
		return nil
	} else {
		return this.properties.GetProperty(aName).Value
	}
}
func (this *Framework) SetProperty(aName string, aValue interface{}) {
	this.properties.SetProperty(aName,aValue)
}
// Obtain the list of the loaded plugins
func (this *Framework) GetBundles() []Bundle {
	return this.bundles
}
func (this *Framework) GetBundleContext() BundleContext {
	return &bundleContextImpl{bundle : this, framework : this, logger : this.logger}
}
// Load the plugins in the plugin folder
func (this *Framework) loadBundles() {
	// Browse plugin directory
	bundleDirectory, err := os.Open(this.bundleFolder)
	if err != nil {
		this.logger.Println("WARNING Bundle directory doesn't exist")
		// no bundle directory
		return
	}
	this.logger.Println("INFO Loading bundles...")
	info, err := bundleDirectory.Stat()
	if !info.IsDir() {
		this.logger.Println("WARNING Bundles directory is not a directory")
		return
	}
	files, err := bundleDirectory.Readdir(0)
	if err != nil {
		this.logger.Println("WARNING Unable to browse bundle directory", err)
		return
	}
	for _, file := range files {
		this.loadBundle(file)
	}
}
// Load one plugin
func (this *Framework) loadBundle(file os.FileInfo) {
	defer func() {
		if r := recover(); r != nil {
			this.logger.Println("WARNING Failed to initialize bundle", file.Name(), ":", r)
		}
	}()
	this.logger.Println("DEBUG Loading bundle", file.Name(), "...")
	thePlugin, err := plugin.Open("plugins/" + file.Name())
	if err != nil {
		this.logger.Println("WARNING Bundle has no Init method", file.Name())
	} else {
		context := NewBundleContext(this)
		bundle := NewPluginBundle(thePlugin, file.Name(), context)
		context.(*bundleContextImpl).setBundle(bundle)
		this.bundles = append(this.bundles,bundle)
		this.logger.Println("INFO Bundle",bundle.GetBundleId(),"-", bundle.GetSymbolicName(), "(",bundle.GetVersion(),") loaded")
	}
}
func (this *Framework) RegisterService(aName string, aService interface{}) {
	if this.services[aName] != nil {
		this.logger.Println("ERROR Service "+aName+" allready registered")
	} else {
		this.logger.Println("INFO Service "+aName+" registered")
		this.services[aName] = NewServiceRef(aName,aService)
	}
}
func (this *Framework) GetService(aName string) *ServiceRef {
	return this.services[aName]
}
func (this *Framework) GetServices() []*ServiceRef {
	result := []*ServiceRef{}
	for _,v := range this.services {
		result = append(result,v)
	}
	return result
}