// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCQuotaFailureViolation stackpath rpc quota failure violation
//
// swagger:model stackpath.rpc.QuotaFailure.Violation
type StackpathRPCQuotaFailureViolation struct {

	// description
	Description string `json:"description,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this stackpath rpc quota failure violation
func (m *StackpathRPCQuotaFailureViolation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCQuotaFailureViolation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCQuotaFailureViolation) UnmarshalBinary(b []byte) error {
	var res StackpathRPCQuotaFailureViolation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
