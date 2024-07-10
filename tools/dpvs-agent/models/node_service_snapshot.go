// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeServiceSnapshot node service snapshot
//
// swagger:model NodeServiceSnapshot
type NodeServiceSnapshot struct {

	// node spec
	NodeSpec *DpvsNodeSpec `json:"NodeSpec,omitempty"`

	// services
	Services *VirtualServerList `json:"Services,omitempty"`
}

// Validate validates this node service snapshot
func (m *NodeServiceSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeServiceSnapshot) validateNodeSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeSpec) { // not required
		return nil
	}

	if m.NodeSpec != nil {
		if err := m.NodeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodeSpec")
			}
			return err
		}
	}

	return nil
}

func (m *NodeServiceSnapshot) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	if m.Services != nil {
		if err := m.Services.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Services")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Services")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this node service snapshot based on the context it is used
func (m *NodeServiceSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeServiceSnapshot) contextValidateNodeSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeSpec != nil {
		if err := m.NodeSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodeSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodeSpec")
			}
			return err
		}
	}

	return nil
}

func (m *NodeServiceSnapshot) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	if m.Services != nil {
		if err := m.Services.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Services")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Services")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeServiceSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeServiceSnapshot) UnmarshalBinary(b []byte) error {
	var res NodeServiceSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
