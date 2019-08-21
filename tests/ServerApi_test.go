package tests

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/wangjiaming0909/GoP/server"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func Test_RegisterHandler(t *testing.T) {
	dispatcher := server.NewDispatcher()
	err := dispatcher.RegisterHandler("/", handler, true)
	if err != nil {
		t.Error("error: ", err.Error())
	}
	if dispatcher.Len() != 1 {
		t.Error("len should be: 1, actual: ", dispatcher.Len())
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func BWithParams(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	for p := range ps {
		fmt.Fprint(w, p)
	}
}

func Test_router(t *testing.T) {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/a", Index)
	r.GET("/a/", Index)
	r.GET("/b?a=1", BWithParams)

	log.Fatal(http.ListenAndServe(":10001", r))
}
