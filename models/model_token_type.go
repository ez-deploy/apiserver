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

// ModelTokenType model token type
//
// swagger:model modelTokenType
type ModelTokenType string

func NewModelTokenType(value ModelTokenType) *ModelTokenType {
	v := value
	return &v
}

const (

	// ModelTokenTypeSession captures enum value "session"
	ModelTokenTypeSession ModelTokenType = "session"

	// ModelTokenTypePrivate captures enum value "private"
	ModelTokenTypePrivate ModelTokenType = "private"

	// ModelTokenTypePublic captures enum value "public"
	ModelTokenTypePublic ModelTokenType = "public"
)

// for schema
var modelTokenTypeEnum []interface{}

func init() {
	var res []ModelTokenType
	if err := json.Unmarshal([]byte(`["session","private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelTokenTypeEnum = append(modelTokenTypeEnum, v)
	}
}

func (m ModelTokenType) validateModelTokenTypeEnum(path, location string, value ModelTokenType) error {
	if err := validate.EnumCase(path, location, value, modelTokenTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this model token type
func (m ModelTokenType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModelTokenTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this model token type based on context it is used
func (m ModelTokenType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
