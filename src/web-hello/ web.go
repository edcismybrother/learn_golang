package main

import (
	"net/http"
	"os"
)

func myFunc(w http.ResponseWriter,req *http.Request) {
	path := req.URL.Path[1:]
	f,err := os.Open(path)
	if err == nil {
		b := []byte{}
		f.Read(b)
		w.Write(b)
	}else {
		w.Write([]byte("404,not find this path or file!"))
	}
//	w.Write([]byte(path))
}

func main() {
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}

