package osgi

import (
	"log"
	"github.com/milak/tools/data"
)

// The OSGI context class
type BundleContext struct {
	bundle 			*Bundle
	framework 		*framework
	Logger			*log.Logger
}
func NewBundleContext(aBundle *Bundle, aFramework *framework) *BundleContext {
	result := BundleContext{bundle : aBundle, framework : aFramework}
	result.Logger = aFramework.Logger
	return &result
}
func (this *BundleContext) GetProperty(aName string) interface{} {
	return framework.GerProperty(aName)
}