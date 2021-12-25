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

// ProjectListPodsResp project list pods resp
//
// swagger:model projectListPodsResp
type ProjectListPodsResp struct {

	// error
	Error *ModelError `json:"error,omitempty"`

	// pods
	Pods []*ModelPod `json:"pods"`
}

// Validate validates this project list pods resp
func (m *ProjectListPodsResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListPodsResp) validateError(formats strfmt.Registry) error {
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

func (m *ProjectListPodsResp) validatePods(formats strfmt.Registry) error {
	if swag.IsZero(m.Pods) { // not required
		return nil
	}

	for i := 0; i < len(m.Pods); i++ {
		if swag.IsZero(m.Pods[i]) { // not required
			continue
		}

		if m.Pods[i] != nil {
			if err := m.Pods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this project list pods resp based on the context it is used
func (m *ProjectListPodsResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListPodsResp) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProjectListPodsResp) contextValidatePods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Pods); i++ {

		if m.Pods[i] != nil {
			if err := m.Pods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectListPodsResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListPodsResp) UnmarshalBinary(b []byte) error {
	var res ProjectListPodsResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
