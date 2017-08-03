package osgi

import (

)

type BundleActivator interface {
	GetVersion() string
	GetSymbolicName() string
	Start(aBundleContext BundleContext)
	Stop(aBundleContext BundleContext)
}