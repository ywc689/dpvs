// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVsVipPortRsParams creates a new GetVsVipPortRsParams object
// with the default values initialized.
func NewGetVsVipPortRsParams() GetVsVipPortRsParams {

	var (
		// initialize parameters with default values

		snapshotDefault = bool(true)
		statsDefault    = bool(false)
	)

	return GetVsVipPortRsParams{
		Snapshot: &snapshotDefault,

		Stats: &statsDefault,
	}
}

// GetVsVipPortRsParams contains all the bound params for the get vs vip port rs operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetVsVipPortRs
type GetVsVipPortRsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	VipPort string
	/*
	  In: query
	  Default: true
	*/
	Snapshot *bool
	/*
	  In: query
	  Default: false
	*/
	Stats *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVsVipPortRsParams() beforehand.
func (o *GetVsVipPortRsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rVipPort, rhkVipPort, _ := route.Params.GetOK("VipPort")
	if err := o.bindVipPort(rVipPort, rhkVipPort, route.Formats); err != nil {
		res = append(res, err)
	}

	qSnapshot, qhkSnapshot, _ := qs.GetOK("snapshot")
	if err := o.bindSnapshot(qSnapshot, qhkSnapshot, route.Formats); err != nil {
		res = append(res, err)
	}

	qStats, qhkStats, _ := qs.GetOK("stats")
	if err := o.bindStats(qStats, qhkStats, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVipPort binds and validates parameter VipPort from path.
func (o *GetVsVipPortRsParams) bindVipPort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.VipPort = raw

	return nil
}

// bindSnapshot binds and validates parameter Snapshot from query.
func (o *GetVsVipPortRsParams) bindSnapshot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVsVipPortRsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("snapshot", "query", "bool", raw)
	}
	o.Snapshot = &value

	return nil
}

// bindStats binds and validates parameter Stats from query.
func (o *GetVsVipPortRsParams) bindStats(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVsVipPortRsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("stats", "query", "bool", raw)
	}
	o.Stats = &value

	return nil
}
