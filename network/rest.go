package network
import (
	"net/http"
	"strings"
)
var packageName string
func Listen(aRoot string, aPort string, aPackageName string){
   packageName = aPackageName
   http.HandleFunc(aRoot, internalHttpListener)
   http.ListenAndServe(":"+aPort, nil)
}
func root(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Milak rest API"))
}
func internalHttpListener(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if len(path) == 0 {
		root(w,req)
	} else {
		if path[0] == '/' {
			path = path[1:]
		}
		w.Write([]byte(path))
		pos := strings.Index(path, "/")
		var object string
		if pos == -1 {
			object = path[0:pos]
		} else {
			object = path[0:]
		}
	}
}
