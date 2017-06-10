// Some objects or interfaces used for list objects
package list
// Equivalent of Iterator java class
type Iterator interface {
	// Returns true if the iterator has more elements to get
	HasNext() bool
	// Returns the next element or nil if no more element is available
	Next() interface{}
}
