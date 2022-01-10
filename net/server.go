package main

import "net/http"

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello hushengs"
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/v1/f1", f1)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
