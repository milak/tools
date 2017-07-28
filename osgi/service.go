package osgi

type Service interface {
	GetName() string
	Get() interface{}
	Start()
	Stop()
}
type serviceImpl struct {
	name 	string
	object 	interface{}
}
func NewService(aName string, aObject interface{}) Service {
	return &serviceImpl{name : aName, object : aObject}
}