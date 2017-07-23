package event
import (
	"github.com/milak/tools/data"
)
type GenericEvent struct {
	Name string
	Properties []*data.Property
}