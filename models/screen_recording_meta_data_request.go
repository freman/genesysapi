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

// ScreenRecordingMetaDataRequest screen recording meta data request
//
// swagger:model ScreenRecordingMetaDataRequest
type ScreenRecordingMetaDataRequest struct {

	// meta data
	MetaData []*ScreenRecordingMetaData `json:"metaData"`

	// participant jid
	ParticipantJid string `json:"participantJid,omitempty"`

	// room Id
	RoomID string `json:"roomId,omitempty"`
}

// Validate validates this screen recording meta data request
func (m *ScreenRecordingMetaDataRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScreenRecordingMetaDataRequest) validateMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	for i := 0; i < len(m.MetaData); i++ {
		if swag.IsZero(m.MetaData[i]) { // not required
			continue
		}

		if m.MetaData[i] != nil {
			if err := m.MetaData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metaData" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metaData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this screen recording meta data request based on the context it is used
func (m *ScreenRecordingMetaDataRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScreenRecordingMetaDataRequest) contextValidateMetaData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetaData); i++ {

		if m.MetaData[i] != nil {
			if err := m.MetaData[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metaData" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metaData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScreenRecordingMetaDataRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScreenRecordingMetaDataRequest) UnmarshalBinary(b []byte) error {
	var res ScreenRecordingMetaDataRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
