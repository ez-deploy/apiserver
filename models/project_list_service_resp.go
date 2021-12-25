// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectListServiceResp project list service resp
//
// swagger:model projectListServiceResp
type ProjectListServiceResp struct {

	// error
	Error *ModelError `json:"error,omitempty"`

	// service
	Service []*ModelService `json:"service"`
}

// Validate validates this project list service resp
func (m *ProjectListServiceResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListServiceResp) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectListServiceResp) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	for i := 0; i < len(m.Service); i++ {
		if swag.IsZero(m.Service[i]) { // not required
			continue
		}

		if m.Service[i] != nil {
			if err := m.Service[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this project list service resp based on the context it is used
func (m *ProjectListServiceResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListServiceResp) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectListServiceResp) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Service); i++ {

		if m.Service[i] != nil {
			if err := m.Service[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectListServiceResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListServiceResp) UnmarshalBinary(b []byte) error {
	var res ProjectListServiceResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
