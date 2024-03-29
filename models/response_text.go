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

// ResponseText Contains information about the text associated with a response.
//
// swagger:model ResponseText
type ResponseText struct {

	// Response text content.
	// Required: true
	Content *string `json:"content"`

	// Response text content type.
	// Enum: [text/plain text/html]
	ContentType string `json:"contentType,omitempty"`
}

// Validate validates this response text
func (m *ResponseText) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseText) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

var responseTextTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text/plain","text/html"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responseTextTypeContentTypePropEnum = append(responseTextTypeContentTypePropEnum, v)
	}
}

const (

	// ResponseTextContentTypeTextPlain captures enum value "text/plain"
	ResponseTextContentTypeTextPlain string = "text/plain"

	// ResponseTextContentTypeTextHTML captures enum value "text/html"
	ResponseTextContentTypeTextHTML string = "text/html"
)

// prop value enum
func (m *ResponseText) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responseTextTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponseText) validateContentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this response text based on context it is used
func (m *ResponseText) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseText) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseText) UnmarshalBinary(b []byte) error {
	var res ResponseText
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
