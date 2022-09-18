// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerspbRegisterCustomerRequest customerspb register customer request
//
// swagger:model customerspbRegisterCustomerRequest
type CustomerspbRegisterCustomerRequest struct {

	// name
	Name string `json:"name,omitempty"`

	// sms number
	SmsNumber string `json:"smsNumber,omitempty"`
}

// Validate validates this customerspb register customer request
func (m *CustomerspbRegisterCustomerRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this customerspb register customer request based on context it is used
func (m *CustomerspbRegisterCustomerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomerspbRegisterCustomerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerspbRegisterCustomerRequest) UnmarshalBinary(b []byte) error {
	var res CustomerspbRegisterCustomerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}