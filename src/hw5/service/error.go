package service

import(
	"net/http"
)

func NotImplemented(w http.ResponseWriter, r *http.Request){
	http.Error(w, "501 Not Implemented", 501)
}

func NotImplementedHandler() http.Handler{
	return http.HandlerFunc(NotImplemented)
}