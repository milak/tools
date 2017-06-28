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
  // register the client object
  objectMap := make(map[string]interface{})
  objectMap["client"] = &client{}
  // listen
  network.Listen("myApp/API","8080",objectMap)
}
/* */
func ExampleNewRestListener(){
}
