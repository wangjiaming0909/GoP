package main

import (
	"container/list"
	"errors"
	"fmt"
	"net/http"
)

type Handler_t func(w http.ResponseWriter, r *http.Request)

type HttpRequestDispatcher struct {
	dispathcTable_ map[string]Handler_t
}

func (dispatcher *HttpRequestDispatcher) RegisterHandler(path string, handler Handler_t, replace bool) error {
	if dispatcher.dispathcTable_[path] != nil {
		dispatcher.dispathcTable_[path] = handler
		return nil
	}
	if replace {
		dispatcher.dispathcTable_[path] = handler
		return nil
	}
	return errors.New("already existed")
}

func a(l *list.List) {
	l.PushBack(1)
	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Println(item)
	}
}
