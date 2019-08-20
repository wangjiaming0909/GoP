package server

import (
	"errors"
	"net/http"
)

//HandlerT dispatch callback
type HandlerT func(w http.ResponseWriter, r *http.Request)

//HTTPRequestDispatcher dispatcher
type HTTPRequestDispatcher struct {
	dispathcTable map[string]HandlerT
}

// RegisterHandler register
func (dispatcher *HTTPRequestDispatcher) RegisterHandler(path string, handler HandlerT, replace bool) error {
	if dispatcher.dispathcTable == nil {
		return errors.New("null reference")
	}
	if dispatcher.dispathcTable[path] == nil {
		dispatcher.dispathcTable[path] = handler
		return nil
	}
	if replace {
		dispatcher.dispathcTable[path] = handler
		return nil
	}
	return errors.New("already existed")
}

// Dispatch dispatch
func (dispatcher *HTTPRequestDispatcher) Dispatch(path string) error{
	return nil
}

// NewDispatcher newDispatcher
func NewDispatcher() *HTTPRequestDispatcher {
	ret := new(HTTPRequestDispatcher)
	ret.dispathcTable = map[string]HandlerT{}
	return ret
}

// Len number of reigstrered handlers
func (dispatcher *HTTPRequestDispatcher) Len() int {
	return len(dispatcher.dispathcTable)
}