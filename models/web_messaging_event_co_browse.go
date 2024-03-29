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

// WebMessagingEventCoBrowse A Cobrowse event.
//
// swagger:model WebMessagingEventCoBrowse
type WebMessagingEventCoBrowse struct {

	// The Cobrowse session ID.
	SessionID string `json:"sessionId,omitempty"`

	// The Cobrowse session join token.
	SessionJoinToken string `json:"sessionJoinToken,omitempty"`

	// Describes the type of Cobrowse event.
	// Required: true
	// Enum: [Offering OfferingExpired OfferingAccepted OfferingRejected]
	Type *string `json:"type"`
}

// Validate validates this web messaging event co browse
func (m *WebMessagingEventCoBrowse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webMessagingEventCoBrowseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offering","OfferingExpired","OfferingAccepted","OfferingRejected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webMessagingEventCoBrowseTypeTypePropEnum = append(webMessagingEventCoBrowseTypeTypePropEnum, v)
	}
}

const (

	// WebMessagingEventCoBrowseTypeOffering captures enum value "Offering"
	WebMessagingEventCoBrowseTypeOffering string = "Offering"

	// WebMessagingEventCoBrowseTypeOfferingExpired captures enum value "OfferingExpired"
	WebMessagingEventCoBrowseTypeOfferingExpired string = "OfferingExpired"

	// WebMessagingEventCoBrowseTypeOfferingAccepted captures enum value "OfferingAccepted"
	WebMessagingEventCoBrowseTypeOfferingAccepted string = "OfferingAccepted"

	// WebMessagingEventCoBrowseTypeOfferingRejected captures enum value "OfferingRejected"
	WebMessagingEventCoBrowseTypeOfferingRejected string = "OfferingRejected"
)

// prop value enum
func (m *WebMessagingEventCoBrowse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webMessagingEventCoBrowseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebMessagingEventCoBrowse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web messaging event co browse based on context it is used
func (m *WebMessagingEventCoBrowse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingEventCoBrowse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingEventCoBrowse) UnmarshalBinary(b []byte) error {
	var res WebMessagingEventCoBrowse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
