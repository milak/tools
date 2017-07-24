package osgi

import (
	"log"
)

type BundleContext interface {
	GetLogger() *log.Logger
	GetProperty(aName string) interface{}
}

// The OSGI context class
type bundleContextImpl struct {
	bundle 			Bundle
	framework 		*Framework
	logger			*log.Logger
}
func NewBundleContext(aFramework *Framework) BundleContext {
	result := BundleContextImpl{framework : aFramework}
	result.Logger = aFramework.Logger
	return &result
}
func (this *bundleContextImpl) GetLogger() *log.Logger {
	return this.logger
}
func (this *bundleContextImpl) setBundle(aBundle Bundle){
	this.bundle = aBundle
}
func (this *bundleContextImpl) GetProperty(aName string) interface{} {
	return this.framework.GetProperty(aName)
}