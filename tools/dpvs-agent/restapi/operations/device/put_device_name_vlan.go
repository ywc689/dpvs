// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutDeviceNameVlanHandlerFunc turns a function with the right signature into a put device name vlan handler
type PutDeviceNameVlanHandlerFunc func(PutDeviceNameVlanParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutDeviceNameVlanHandlerFunc) Handle(params PutDeviceNameVlanParams) middleware.Responder {
	return fn(params)
}

// PutDeviceNameVlanHandler interface for that can handle valid put device name vlan params
type PutDeviceNameVlanHandler interface {
	Handle(PutDeviceNameVlanParams) middleware.Responder
}

// NewPutDeviceNameVlan creates a new http.Handler for the put device name vlan operation
func NewPutDeviceNameVlan(ctx *middleware.Context, handler PutDeviceNameVlanHandler) *PutDeviceNameVlan {
	return &PutDeviceNameVlan{Context: ctx, Handler: handler}
}

/*
	PutDeviceNameVlan swagger:route PUT /device/{name}/vlan device putDeviceNameVlan

add/update special net device
*/
type PutDeviceNameVlan struct {
	Context *middleware.Context
	Handler PutDeviceNameVlanHandler
}

func (o *PutDeviceNameVlan) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutDeviceNameVlanParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
