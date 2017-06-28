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
Listen on port 'aPort' and the context root 'aRoot'. The map given as argument is used to call objects according the url.
Example : objectMap["client"] = &MyClientObject{}
Where "client" is the object used in th URL : http://myApp/API/client and MyClientObject is the target object called.
This method i a all-in-one method. For specific use, you can use NewRestListener() method.
*/
func Listen(aRoot string, aPort string, aObjectMap map[string]interface{}){
	if !strings.HasSuffix(aRoot, "/") {
		aRoot = aRoot + "/"
	}
	http.HandleFunc(aRoot, NewRestListener(aRoot,aObjectMap))
	http.ListenAndServe(":"+aPort, nil)
}

/* Internal object that will be used for listening http calls */
type restListener struct {
	contextRoot string
	objectMap map[string]interface{}
}
/* The listener called by the HttpListener */
func (this *restListener) internalHttpListener(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	path = path[len(this.contextRoot):]
	if len(path) == 0 || path == "/" {
		this.root(w,req,"")
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
		object := this.objectMap[objectName]
		if object == nil {
			this.root(w,req,"Object not found " + objectName)
		} else {
			if req.Method == http.MethodGet {
				theObject, ok := object.(get)
				if ok {
				    theObject.Get(w, req)
				} else {
					this.root(w,req,"Get method not found on " + objectName)
				}
			} else if req.Method == http.MethodPost {
				theObject, ok := object.(post)
				if ok {
				    theObject.Post(w, req)
				} else {
					this.root(w,req,"Post method not found on " + objectName)
				}
			} else if req.Method == http.MethodDelete {
				theObject, ok := object.(del)
				if ok {
				    theObject.Delete(w, req)
				} else {
					this.root(w,req,"Delete method not found on " + objectName)
				}
			} else if req.Method == http.MethodPut {
				theObject, ok := object.(put)
				if ok {
				    theObject.Put(w, req)
				} else {
					this.root(w,req,"Put method not found on " + objectName)
				}
			} else {
				this.root(w,req,"Unsupported method " + req.Method)
			}
		}
	}
}
/* This method displays the root page */
func (this *restListener) root(w http.ResponseWriter, req *http.Request, aError string){
	w.Write([]byte("<html><body>"))
	w.Write([]byte("<h1>Milak rest API</h1>"))
	if len(aError) != 0 {
		w.Write([]byte("Error : <span style='color:red'>"+aError+"</span><br/>"))
	}
	w.Write([]byte("Available objects : <br/>"))
	w.Write([]byte("<table style='width:100%;border:solid 1px'>"))
	w.Write([]byte("<thead style='background-color:#9090F0'><tr><th>object</th><th>get</th><th>post</th><th>put</th><th>delete</th></tr></thead>"))
	w.Write([]byte("<tbody>"))
	for k,o := range this.objectMap {
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
/*
For specific use, you could need to use NewRestListener insteed Listen() method.
*/
func NewRestListener(aRoot string, aObjectMap map[string]interface{}) (func(w http.ResponseWriter, req *http.Request)) {
	restListener := &restListener{contextRoot : aRoot, objectMap : aObjectMap}
	return restListener.internalHttpListener
}
