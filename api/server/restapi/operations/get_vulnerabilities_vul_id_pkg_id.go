// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetVulnerabilitiesVulIDPkgIDHandlerFunc turns a function with the right signature into a get vulnerabilities vul ID pkg ID handler
type GetVulnerabilitiesVulIDPkgIDHandlerFunc func(GetVulnerabilitiesVulIDPkgIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVulnerabilitiesVulIDPkgIDHandlerFunc) Handle(params GetVulnerabilitiesVulIDPkgIDParams) middleware.Responder {
	return fn(params)
}

// GetVulnerabilitiesVulIDPkgIDHandler interface for that can handle valid get vulnerabilities vul ID pkg ID params
type GetVulnerabilitiesVulIDPkgIDHandler interface {
	Handle(GetVulnerabilitiesVulIDPkgIDParams) middleware.Responder
}

// NewGetVulnerabilitiesVulIDPkgID creates a new http.Handler for the get vulnerabilities vul ID pkg ID operation
func NewGetVulnerabilitiesVulIDPkgID(ctx *middleware.Context, handler GetVulnerabilitiesVulIDPkgIDHandler) *GetVulnerabilitiesVulIDPkgID {
	return &GetVulnerabilitiesVulIDPkgID{Context: ctx, Handler: handler}
}

/* GetVulnerabilitiesVulIDPkgID swagger:route GET /vulnerabilities/{vul_id}/{pkg_id} getVulnerabilitiesVulIdPkgId

Get a vulnerability of a specific package

*/
type GetVulnerabilitiesVulIDPkgID struct {
	Context *middleware.Context
	Handler GetVulnerabilitiesVulIDPkgIDHandler
}

func (o *GetVulnerabilitiesVulIDPkgID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetVulnerabilitiesVulIDPkgIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
