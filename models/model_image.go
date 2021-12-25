// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelImage model image
//
// swagger:model modelImage
type ModelImage struct {

	// url
	URL string `json:"url,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this model image
func (m *ModelImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model image based on context it is used
func (m *ModelImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelImage) UnmarshalBinary(b []byte) error {
	var res ModelImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
