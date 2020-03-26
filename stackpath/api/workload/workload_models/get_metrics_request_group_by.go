// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GetMetricsRequestGroupBy The available fields to group instance metrics by
//
// swagger:model GetMetricsRequestGroupBy
type GetMetricsRequestGroupBy string

const (

	// GetMetricsRequestGroupByNONE captures enum value "NONE"
	GetMetricsRequestGroupByNONE GetMetricsRequestGroupBy = "NONE"

	// GetMetricsRequestGroupByINSTANCENAME captures enum value "INSTANCE_NAME"
	GetMetricsRequestGroupByINSTANCENAME GetMetricsRequestGroupBy = "INSTANCE_NAME"
)

// for schema
var getMetricsRequestGroupByEnum []interface{}

func init() {
	var res []GetMetricsRequestGroupBy
	if err := json.Unmarshal([]byte(`["NONE","INSTANCE_NAME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getMetricsRequestGroupByEnum = append(getMetricsRequestGroupByEnum, v)
	}
}

func (m GetMetricsRequestGroupBy) validateGetMetricsRequestGroupByEnum(path, location string, value GetMetricsRequestGroupBy) error {
	if err := validate.Enum(path, location, value, getMetricsRequestGroupByEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this get metrics request group by
func (m GetMetricsRequestGroupBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGetMetricsRequestGroupByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
