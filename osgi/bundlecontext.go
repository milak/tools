package osgi

import (
	"log"
)

type BundleContext interface {
	Logger			*log.Logger
	GetProperty(aName string) interface{}
}

// The OSGI context class
type bundleContextImpl struct {
	bundle 			Bundle
	framework 		*Framework
	Logger			*log.Logger
}
func NewBundleContext(aFramework *Framework) BundleContext {
	result := BundleContextImpl{framework : aFramework}
	result.Logger = aFramework.Logger
	return &result
}
func (this *BundleContext) setBundle(aBundle Bundle){
	this.bundle = aBundle
}
func (this *BundleContext) GetProperty(aName string) interface{} {
	return this.framework.GetProperty(aName)
}