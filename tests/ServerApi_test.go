package tests

import (
	"net/http"
	"testing"

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
