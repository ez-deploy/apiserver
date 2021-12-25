// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelPort Port for container.
//
// swagger:model modelPort
type ModelPort struct {

	// name
	Name string `json:"name,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`
}

// Validate validates this model port
func (m *ModelPort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model port based on context it is used
func (m *ModelPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelPort) UnmarshalBinary(b []byte) error {
	var res ModelPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}