package network
func ExampleGetLocalIP() {
  address, err := GetLocalIP()
  if err != nil {
     fmt.Println("Error getting IP address",err)
  } else {
     fmt.Println("IP address is",address)
  }
}
// Declare client object :
type client struct {
}
func (this *client) Get(w http.ResponseWriter, req *http.Request) {
   // process
}
/* Sample for a call of http://host:8080/myApp/API/client */
func ExampleListen() {
  // Add specific listener
  http.HandleFunc("myApp/", MyListenerFunction)
  
  // Register the client object
  objectMap := make(map[string]interface{})
  objectMap["client"] = &client{}
  
  // Listen
  network.Listen("myApp/API","8080",objectMap)
}
/* Sample for a call of http://host:8080/myApp/API/client */
func ExampleNewRestListener(){
  // Add specific listener
  http.HandleFunc("myApp/", MyListenerFunction)
  
  // Register the client object
  objectMap := make(map[string]interface{})
  objectMap["client"] = &client{}
  // Add RestListener
  http.HandleFunc("myApp/API/", NewRestListener("myApp/API/",aObjectMap))
  
  // Listen
  http.ListenAndServe(":8080", nil)
}
