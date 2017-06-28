package network
import (
	"net/http"
	"strings"
)
/* Interface of an object supporting GET method */
type get interface {
	Get(w http.ResponseWriter, req *http.Request)
}
/* Interface of an object supporting POST method */
type post interface {
	Post(w http.ResponseWriter, req *http.Request)
}
/* Interface of an object supporting DEL method */
type del interface {
	Delete(w http.ResponseWriter, req *http.Request)
}
/* Interface of an object supporting PUT method */
type put interface {
	Put(w http.ResponseWriter, req *http.Request)
}
/*
Listen on port aPort and the context root aRoot. The map given as argument is used to call objects according the url.

Example :

for a call of http://host:8080/myApp/API/client
  // Declare client object :
  type client struct {
  }
  func (this *client) Get(w http.ResponseWriter, req *http.Request) {
     // process
  }
  // register the client object
  objectMap := make(map[string]interface{})
  objectMap["client"] = $client{}
  // listen
  network.Listen("myApp/API","8080",objectMap)
*/
func Listen(aRoot string, aPort string, aObjectMap map[string]interface{}){
	if !strings.HasSuffix(aRoot, "/") {
		aRoot = aRoot + "/"
	}
	http.HandleFunc(aRoot, NewRestListener(aRoot,aObjectMap))
	http.ListenAndServe(":"+aPort, nil)
}
/* This method displays the root page */
func root(w http.ResponseWriter, req *http.Request, aError string){
	w.Write([]byte("<html><body>"))
	w.Write([]byte("<h1>Milak rest API</h1>"))
	if len(aError) != 0 {
		w.Write([]byte("Error : <span style='color:red'>"+aError+"</span><br/>"))
	}
	w.Write([]byte("Available objects : <br/>"))
	w.Write([]byte("<table style='width:100%;border:solid 1px'>"))
	w.Write([]byte("<thead style='background-color:#9090F0'><tr><th>object</th><th>get</th><th>post</th><th>put</th><th>delete</th></tr></thead>"))
	w.Write([]byte("<tbody>"))
	for k,o := range objectMap {
		w.Write([]byte("<tr><td>"+k+"</td>"))
		_, ok := o.(get)
		if ok {
		    w.Write([]byte("<td><a href='"+k+"'>call</a></td>"))
		} else {
			w.Write([]byte("<td><i>---</i></td>"))
		}
		_, ok = o.(post)
		if ok {
		    w.Write([]byte("<td><form action='"+k+"' method='POST'><input type='submit'></input></form></td>"))
		} else {
			w.Write([]byte("<td><i>---</i></td>"))
		}
		_, ok = o.(put)
		if ok {
		    w.Write([]byte("<td><form action='"+k+"' method='PUT'><input type='submit'></input></form></td>"))
		} else {
			w.Write([]byte("<td><i>---</i></td>"))
		}
		_, ok = o.(del)
		if ok {
		    w.Write([]byte("<td><form action='"+k+"' method='DELETE'><input type='submit'></input></form></td></tr>"))
		} else {
			w.Write([]byte("<td><i>---</i></td></tr>"))
		}
	}
	w.Write([]byte("</tbody>"))
	w.Write([]byte("</table>"))
	w.Write([]byte("</body></html>"))
}
type restListener struct {
	contextRoot string
	objectMap map[string]interface{}
}
/* The listener called by the HttpListener */
func (this *restListener) internalHttpListener(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	path = path[len(contextRoot):]
	if len(path) == 0 || path == "/" {
		root(w,req,"")
	} else {
		if path[0] == '/' {
			path = path[1:]
		}
		pos := strings.Index(path, "/")
		var objectName string
		if pos == -1 {
			objectName = path[0:]
		} else {
			objectName = path[0:pos]
		}
		object := objectMap[objectName]
		if object == nil {
			root(w,req,"Object not found " + objectName)
		} else {
			if req.Method == http.MethodGet {
				theObject, ok := object.(get)
				if ok {
				    theObject.Get(w, req)
				} else {
					root(w,req,"Get method not found on " + objectName)
				}
			} else if req.Method == http.MethodPost {
				theObject, ok := object.(post)
				if ok {
				    theObject.Post(w, req)
				} else {
					root(w,req,"Post method not found on " + objectName)
				}
			} else if req.Method == http.MethodDelete {
				theObject, ok := object.(del)
				if ok {
				    theObject.Delete(w, req)
				} else {
					root(w,req,"Delete method not found on " + objectName)
				}
			} else if req.Method == http.MethodPut {
				theObject, ok := object.(put)
				if ok {
				    theObject.Put(w, req)
				} else {
					root(w,req,"Put method not found on " + objectName)
				}
			} else {
				root(w,req,"Unsupported method " + req.Method)
			}
		}
	}
}
/*
 For specific use, you could need to use
*/
func NewRestListener(aRoot string, aObjectMap map[string]interface{}){
	var restListener := &restListener{contextRoot : aRoot, objectMap : aObjectMap}
	return restListener.internalHttpListener
}
