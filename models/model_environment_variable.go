// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelEnvironmentVariable model environment variable
//
// swagger:model modelEnvironmentVariable
type ModelEnvironmentVariable struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this model environment variable
func (m *ModelEnvironmentVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model environment variable based on context it is used
func (m *ModelEnvironmentVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelEnvironmentVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelEnvironmentVariable) UnmarshalBinary(b []byte) error {
	var res ModelEnvironmentVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
