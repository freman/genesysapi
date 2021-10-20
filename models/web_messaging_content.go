// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebMessagingContent Message content element.
//
// swagger:model WebMessagingContent
type WebMessagingContent struct {

	// Attachment content.
	// Read Only: true
	Attachment *WebMessagingAttachment `json:"attachment,omitempty"`

	// Button response content.
	ButtonResponse *WebMessagingButtonResponse `json:"buttonResponse,omitempty"`

	// Type of this content element. If contentType = "Attachment" only one item is allowed.
	// Read Only: true
	// Enum: [Attachment QuickReply ButtonResponse GenericTemplate]
	ContentType string `json:"contentType,omitempty"`

	// Generic content.
	Generic *WebMessagingGeneric `json:"generic,omitempty"`

	// Quick reply content.
	QuickReply *WebMessagingQuickReply `json:"quickReply,omitempty"`
}

// Validate validates this web messaging content
func (m *WebMessagingContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateButtonResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuickReply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebMessagingContent) validateAttachment(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachment) { // not required
		return nil
	}

	if m.Attachment != nil {
		if err := m.Attachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

func (m *WebMessagingContent) validateButtonResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.ButtonResponse) { // not required
		return nil
	}

	if m.ButtonResponse != nil {
		if err := m.ButtonResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buttonResponse")
			}
			return err
		}
	}

	return nil
}

var webMessagingContentTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attachment","QuickReply","ButtonResponse","GenericTemplate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webMessagingContentTypeContentTypePropEnum = append(webMessagingContentTypeContentTypePropEnum, v)
	}
}

const (

	// WebMessagingContentContentTypeAttachment captures enum value "Attachment"
	WebMessagingContentContentTypeAttachment string = "Attachment"

	// WebMessagingContentContentTypeQuickReply captures enum value "QuickReply"
	WebMessagingContentContentTypeQuickReply string = "QuickReply"

	// WebMessagingContentContentTypeButtonResponse captures enum value "ButtonResponse"
	WebMessagingContentContentTypeButtonResponse string = "ButtonResponse"

	// WebMessagingContentContentTypeGenericTemplate captures enum value "GenericTemplate"
	WebMessagingContentContentTypeGenericTemplate string = "GenericTemplate"
)

// prop value enum
func (m *WebMessagingContent) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webMessagingContentTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebMessagingContent) validateContentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *WebMessagingContent) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(m.Generic) { // not required
		return nil
	}

	if m.Generic != nil {
		if err := m.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generic")
			}
			return err
		}
	}

	return nil
}

func (m *WebMessagingContent) validateQuickReply(formats strfmt.Registry) error {

	if swag.IsZero(m.QuickReply) { // not required
		return nil
	}

	if m.QuickReply != nil {
		if err := m.QuickReply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quickReply")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingContent) UnmarshalBinary(b []byte) error {
	var res WebMessagingContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
