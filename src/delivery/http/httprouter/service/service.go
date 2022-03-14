package service

import "github.com/julienschmidt/httprouter"

type HttpService interface {
	GetHandles() []Handle
}

type Handle struct {
	Method string
	Path   string
	Handle httprouter.Handle
}
