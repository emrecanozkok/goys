package api

import (
	"main/api/controller"
	"net/http"
)

/*
	Redirect to controller.Need to be manupulate in the future
*/

func Set(res http.ResponseWriter, req *http.Request) {
	controller.Set(res, req)
}

func Get(res http.ResponseWriter, req *http.Request) {
	controller.Get(res, req)
}

func Flush(res http.ResponseWriter, req *http.Request) {
	controller.Flush(res, req)
}
