package osgi

type ServiceRef struct {
	name 	string
	object 	interface{}
}
func (this *ServiceRef) GetName() interface{} {
	return this.name
}
func (this *ServiceRef) Get() interface{} {
	return this.object
}
func NewServiceRef(aName string, aObject interface{}) *ServiceRef {
	return &ServiceRef{name : aName, object : aObject}
}