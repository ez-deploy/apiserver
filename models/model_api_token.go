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

// ModelAPIToken model API token
//
// swagger:model modelAPIToken
type ModelAPIToken struct {

	// allowed actions
	AllowedActions []*ModelActions `json:"allowed_actions"`

	// deadline timestamp
	DeadlineTimestamp string `json:"deadline_timestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// token
	Token *ModelToken `json:"token,omitempty"`
}

// Validate validates this model API token
func (m *ModelAPIToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelAPIToken) validateAllowedActions(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowedActions) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedActions); i++ {
		if swag.IsZero(m.AllowedActions[i]) { // not required
			continue
		}

		if m.AllowedActions[i] != nil {
			if err := m.AllowedActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowed_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowed_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelAPIToken) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this model API token based on the context it is used
func (m *ModelAPIToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelAPIToken) contextValidateAllowedActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedActions); i++ {

		if m.AllowedActions[i] != nil {
			if err := m.AllowedActions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowed_actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowed_actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelAPIToken) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {
		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelAPIToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAPIToken) UnmarshalBinary(b []byte) error {
	var res ModelAPIToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
