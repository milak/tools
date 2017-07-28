package osgi

import (
	"log"
)

type BundleContext interface {
	GetLogger() *log.Logger
	GetBundle() Bundle
	GetProperty(aName string) interface{}
	SetProperty(aName string, aValue interface{})
	RegisterService(aName string, aService interface{})
	GetService(aName string) Service
}

// The OSGI context class
type bundleContextImpl struct {
	bundle 			Bundle
	framework 		*Framework
	logger			*log.Logger
}
func NewBundleContext(aFramework *Framework) BundleContext {
	result := bundleContextImpl{framework : aFramework}
	result.logger = aFramework.Logger
	return &result
}
func (this *bundleContextImpl) GetLogger() *log.Logger {
	return this.logger
}
func (this *bundleContextImpl) setBundle(aBundle Bundle){
	this.bundle = aBundle
}
func (this *bundleContextImpl) GetBundle() Bundle {
	return this.bundle
}
func (this *bundleContextImpl) GetProperty(aName string) interface{} {
	return this.framework.GetProperty(aName)
}
func (this *bundleContextImpl) SetProperty(aName string, aValue interface{}) {
	this.framework.SetProperty(aName, aValue)
}
func (this *bundleContextImpl) RegisterService(aName string, aService interface{}) {
	this.framework.RegisterService(aName,aService)
}
func (this *bundleContextImpl) GetService(aName string) Service {
	return this.framework.GetService(aName)
}