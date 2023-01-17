// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ArchiveRetention archive retention
//
// swagger:model ArchiveRetention
type ArchiveRetention struct {

	// days
	Days int32 `json:"days,omitempty"`

	// storage medium
	// Enum: [CLOUDARCHIVE]
	StorageMedium string `json:"storageMedium,omitempty"`
}

// Validate validates this archive retention
func (m *ArchiveRetention) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageMedium(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var archiveRetentionTypeStorageMediumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLOUDARCHIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		archiveRetentionTypeStorageMediumPropEnum = append(archiveRetentionTypeStorageMediumPropEnum, v)
	}
}

const (

	// ArchiveRetentionStorageMediumCLOUDARCHIVE captures enum value "CLOUDARCHIVE"
	ArchiveRetentionStorageMediumCLOUDARCHIVE string = "CLOUDARCHIVE"
)

// prop value enum
func (m *ArchiveRetention) validateStorageMediumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, archiveRetentionTypeStorageMediumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchiveRetention) validateStorageMedium(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageMedium) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorageMediumEnum("storageMedium", "body", m.StorageMedium); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this archive retention based on context it is used
func (m *ArchiveRetention) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArchiveRetention) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchiveRetention) UnmarshalBinary(b []byte) error {
	var res ArchiveRetention
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
