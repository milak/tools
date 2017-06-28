package network
func ExampleGetLocalIP() {
  address, err := GetLocalIP()
  if err != nil {
     fmt.Println("Error getting IP address",err)
  } else {
     fmt.Println("IP address is",address)
  }
}
