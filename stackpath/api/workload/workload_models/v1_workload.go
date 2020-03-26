// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Workload A computing workload
//
// Workloads define a computing application to deploy to StackPath's edge network.
//
// swagger:model v1Workload
type V1Workload struct {

	// A workload's unique identifier
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *V1Metadata `json:"metadata,omitempty"`

	// A workload's name as it appears in the StackPath portal
	Name string `json:"name,omitempty"`

	// A workload's programmatic name
	//
	// Workload slugs are used to build its instances names
	Slug string `json:"slug,omitempty"`

	// spec
	Spec *V1WorkloadSpec `json:"spec,omitempty"`

	// The ID of the stack that a workload belongs to
	// Read Only: true
	StackID string `json:"stackId,omitempty"`

	// status
	// Read Only: true
	Status V1WorkloadStatus `json:"status,omitempty"`

	// targets
	Targets V1TargetMapEntry `json:"targets,omitempty"`
}

// Validate validates this v1 workload
func (m *V1Workload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Workload) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1Workload) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *V1Workload) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *V1Workload) validateTargets(formats strfmt.Registry) error {

	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	if err := m.Targets.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targets")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Workload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Workload) UnmarshalBinary(b []byte) error {
	var res V1Workload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
