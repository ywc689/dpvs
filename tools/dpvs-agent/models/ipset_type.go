// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IpsetType ipset type
//
// swagger:model IpsetType
type IpsetType string

func NewIpsetType(value IpsetType) *IpsetType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IpsetType.
func (m IpsetType) Pointer() *IpsetType {
	return &m
}

const (

	// IpsetTypeBitmapIP captures enum value "bitmap:ip"
	IpsetTypeBitmapIP IpsetType = "bitmap:ip"

	// IpsetTypeBitmapIPMac captures enum value "bitmap:ip,mac"
	IpsetTypeBitmapIPMac IpsetType = "bitmap:ip,mac"

	// IpsetTypeBitmapPort captures enum value "bitmap:port"
	IpsetTypeBitmapPort IpsetType = "bitmap:port"

	// IpsetTypeHashIP captures enum value "hash:ip"
	IpsetTypeHashIP IpsetType = "hash:ip"

	// IpsetTypeHashNet captures enum value "hash:net"
	IpsetTypeHashNet IpsetType = "hash:net"

	// IpsetTypeHashIPPort captures enum value "hash:ip,port"
	IpsetTypeHashIPPort IpsetType = "hash:ip,port"

	// IpsetTypeHashNetPort captures enum value "hash:net,port"
	IpsetTypeHashNetPort IpsetType = "hash:net,port"

	// IpsetTypeHashNetPortIface captures enum value "hash:net,port,iface"
	IpsetTypeHashNetPortIface IpsetType = "hash:net,port,iface"

	// IpsetTypeHashIPPortIP captures enum value "hash:ip,port,ip"
	IpsetTypeHashIPPortIP IpsetType = "hash:ip,port,ip"

	// IpsetTypeHashIPPortNet captures enum value "hash:ip,port,net"
	IpsetTypeHashIPPortNet IpsetType = "hash:ip,port,net"

	// IpsetTypeHashNetPortNet captures enum value "hash:net,port,net"
	IpsetTypeHashNetPortNet IpsetType = "hash:net,port,net"

	// IpsetTypeHashNetPortNetPort captures enum value "hash:net,port,net,port"
	IpsetTypeHashNetPortNetPort IpsetType = "hash:net,port,net,port"
)

// for schema
var ipsetTypeEnum []interface{}

func init() {
	var res []IpsetType
	if err := json.Unmarshal([]byte(`["bitmap:ip","bitmap:ip,mac","bitmap:port","hash:ip","hash:net","hash:ip,port","hash:net,port","hash:net,port,iface","hash:ip,port,ip","hash:ip,port,net","hash:net,port,net","hash:net,port,net,port"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipsetTypeEnum = append(ipsetTypeEnum, v)
	}
}

func (m IpsetType) validateIpsetTypeEnum(path, location string, value IpsetType) error {
	if err := validate.EnumCase(path, location, value, ipsetTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ipset type
func (m IpsetType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIpsetTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ipset type based on context it is used
func (m IpsetType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}