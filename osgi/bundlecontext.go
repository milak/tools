package osgi

import (
	"log"
)

// The OSGI context class
type BundleContext struct {
	bundle 			Bundle
	framework 		*Framework
	Logger			*log.Logger
}
func NewBundleContext(aFramework *Framework) *BundleContext {
	result := BundleContext{framework : aFramework}
	result.Logger = aFramework.Logger
	return &result
}
func (this *BundleContext) setBundle(aBundle Bundle){
	this.bundle = aBundle
}
func (this *BundleContext) GetProperty(aName string) interface{} {
	return this.framework.GetProperty(aName)
}