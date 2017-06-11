package network
var packageName string
func Listen(aPort string,aPackageName string){
   packageName = aPackageName
   http.HandleFunc("/", _internalHttpListener)
   go http.ListenAndServe(":"+this.port, nil)
}
func _internalHttpListener(w http.ResponseWriter, req *http.Request) {{
   w.Write([]byte("Hello !!!"))
}
