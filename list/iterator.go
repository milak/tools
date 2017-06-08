package list

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
