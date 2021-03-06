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

// FaxSendRequest fax send request
//
// swagger:model FaxSendRequest
type FaxSendRequest struct {

	// A list of outbound fax dialing addresses. E.g. +13175555555 or 3175555555
	// Required: true
	Addresses []string `json:"addresses"`

	// The content type that is going to be uploaded. If Content Management document is used for faxing, contentType will be ignored
	// Enum: [application/pdf image/tiff application/msword application/vnd.oasis.opendocument.text application/vnd.openxmlformats-officedocument.wordprocessingml.document]
	ContentType string `json:"contentType,omitempty"`

	// Data for coversheet generation.
	CoverSheet *CoverSheet `json:"coverSheet,omitempty"`

	// DocumentId of Content Management artifact. If Content Management document is not used for faxing, documentId should be null
	DocumentID string `json:"documentId,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Time zone offset minutes from GMT
	TimeZoneOffsetMinutes int32 `json:"timeZoneOffsetMinutes,omitempty"`

	// Workspace in which the document should be stored. If Content Management document is used for faxing, workspace will be ignored
	Workspace *Workspace `json:"workspace,omitempty"`
}

// Validate validates this fax send request
func (m *FaxSendRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverSheet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FaxSendRequest) validateAddresses(formats strfmt.Registry) error {

	if err := validate.Required("addresses", "body", m.Addresses); err != nil {
		return err
	}

	return nil
}

var faxSendRequestTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["application/pdf","image/tiff","application/msword","application/vnd.oasis.opendocument.text","application/vnd.openxmlformats-officedocument.wordprocessingml.document"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		faxSendRequestTypeContentTypePropEnum = append(faxSendRequestTypeContentTypePropEnum, v)
	}
}

const (

	// FaxSendRequestContentTypeApplicationPdf captures enum value "application/pdf"
	FaxSendRequestContentTypeApplicationPdf string = "application/pdf"

	// FaxSendRequestContentTypeImageTiff captures enum value "image/tiff"
	FaxSendRequestContentTypeImageTiff string = "image/tiff"

	// FaxSendRequestContentTypeApplicationMsword captures enum value "application/msword"
	FaxSendRequestContentTypeApplicationMsword string = "application/msword"

	// FaxSendRequestContentTypeApplicationVndOasisOpendocumentText captures enum value "application/vnd.oasis.opendocument.text"
	FaxSendRequestContentTypeApplicationVndOasisOpendocumentText string = "application/vnd.oasis.opendocument.text"

	// FaxSendRequestContentTypeApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument captures enum value "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	FaxSendRequestContentTypeApplicationVndOpenxmlformatsOfficedocumentWordprocessingmlDocument string = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
)

// prop value enum
func (m *FaxSendRequest) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, faxSendRequestTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FaxSendRequest) validateContentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *FaxSendRequest) validateCoverSheet(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverSheet) { // not required
		return nil
	}

	if m.CoverSheet != nil {
		if err := m.CoverSheet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coverSheet")
			}
			return err
		}
	}

	return nil
}

func (m *FaxSendRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxSendRequest) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FaxSendRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FaxSendRequest) UnmarshalBinary(b []byte) error {
	var res FaxSendRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
