// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GetNetworkPolicyResponse A response from a request to get a network policy to a stack
//
// swagger:model v1GetNetworkPolicyResponse
type V1GetNetworkPolicyResponse struct {

	// network policy
	NetworkPolicy *V1NetworkPolicy `json:"networkPolicy,omitempty"`
}

// Validate validates this v1 get network policy response
func (m *V1GetNetworkPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetNetworkPolicyResponse) validateNetworkPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkPolicy) { // not required
		return nil
	}

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetNetworkPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetNetworkPolicyResponse) UnmarshalBinary(b []byte) error {
	var res V1GetNetworkPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
