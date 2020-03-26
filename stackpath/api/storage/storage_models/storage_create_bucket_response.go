// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageCreateBucketResponse The bucket created
//
// swagger:model storageCreateBucketResponse
type StorageCreateBucketResponse struct {

	// bucket
	Bucket *StorageBucket `json:"bucket,omitempty"`
}

// Validate validates this storage create bucket response
func (m *StorageCreateBucketResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageCreateBucketResponse) validateBucket(formats strfmt.Registry) error {

	if swag.IsZero(m.Bucket) { // not required
		return nil
	}

	if m.Bucket != nil {
		if err := m.Bucket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bucket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageCreateBucketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCreateBucketResponse) UnmarshalBinary(b []byte) error {
	var res StorageCreateBucketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
