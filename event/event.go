// A very simple event bus
// Simple event bus that allows to register as a listener.
// How to use :
// 
// 1 - define an event :
//   type MyEvent struct {
//     MyInfo string
//   }
// 2 - define a listener
//   type MyListener struct {
//   }
//   func (this *MyListener) Event(aEvent interface{}) {
//     switch e := aEvent.(type) {
//         case conf.MyEvent :
//         fmt.Println(e.MyInfo)
//     }
//   }
// 3 - register
//   var myListener MyListener
//   event.EventBus.AddListener(myListener)
// 4 - fire event
//   event.EventBus.FireEvent(&MyEvent{"Hello world !"})
package event
// The Interface an EventListener has to implement to register to the event bus
type Listener interface {
	Event(aEvent interface{})
}
// A global instance of EventBus
var EventBus eventBus
type eventBus struct {
	listeners		[]Listener
}
// Register a listener for events
func (this *eventBus) AddListener(aListener Listener) {
	this.listeners = append(this.listeners,aListener)
}
// Remove a previously registered listener
func (this *eventBus) RemoveListener(aListener Listener) {
	for i,l := range this.listeners {
		if l == aListener {
			this.listeners = append(this.listeners[0:i],this.listeners[i:]...)
			return
		}
	}
}
// Fire an event
// Every listener, is called as a new thread :
// go listener.Event(aEvent)
func (this *eventBus) FireEvent(aEvent interface{}) {
	for _,listener := range this.listeners {
		go listener.Event(aEvent)
	}
}
