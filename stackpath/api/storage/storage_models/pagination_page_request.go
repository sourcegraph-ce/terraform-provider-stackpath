// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationPageRequest Pagination request information
//
// This is modeled after the GraphQL Relay spec to support both cursor based pagination and traditional offset based pagination.
//
// swagger:model paginationPageRequest
type PaginationPageRequest struct {

	// The cursor value after which data will be returned
	After string `json:"after,omitempty"`

	// SQL-style constraint filters
	Filter string `json:"filter,omitempty"`

	// The number of items desired
	First string `json:"first,omitempty"`

	// Sort the response by the given field
	SortBy string `json:"sortBy,omitempty"`
}

// Validate validates this pagination page request
func (m *PaginationPageRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationPageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationPageRequest) UnmarshalBinary(b []byte) error {
	var res PaginationPageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
