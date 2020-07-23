// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DownloadResponse download response
//
// swagger:model DownloadResponse
type DownloadResponse struct {

	// content location Uri
	ContentLocationURI string `json:"contentLocationUri,omitempty"`

	// image Uri
	ImageURI string `json:"imageUri,omitempty"`

	// thumbnails
	Thumbnails []*DocumentThumbnail `json:"thumbnails"`
}

// Validate validates this download response
func (m *DownloadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThumbnails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DownloadResponse) validateThumbnails(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumbnails) { // not required
		return nil
	}

	for i := 0; i < len(m.Thumbnails); i++ {
		if swag.IsZero(m.Thumbnails[i]) { // not required
			continue
		}

		if m.Thumbnails[i] != nil {
			if err := m.Thumbnails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("thumbnails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DownloadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadResponse) UnmarshalBinary(b []byte) error {
	var res DownloadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
