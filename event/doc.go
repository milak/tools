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
//   event.Bus.AddListener(myListener)
// 4 - fire event
//   event.Bus.FireEvent(&MyEvent{"Hello world !"})
package event