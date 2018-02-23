package main

import (
	"net/http"
	"io/ioutil"
)

func myFunc(w http.ResponseWriter,req *http.Request) {
	path := req.URL.Path[1:]
	b,e := ioutil.ReadFile(path)
	if e == nil {
		w.Write(b)
	}else{
		w.Write([]byte("404,not find this path or file!"))
	}
//	w.Write([]byte(path))
}

func main() {
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}

