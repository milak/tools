package osgi

type Service interface {
	GetName() string
	Start()
	Stop()
}