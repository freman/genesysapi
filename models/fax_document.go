// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FaxDocument fax document
//
// swagger:model FaxDocument
type FaxDocument struct {

	// caller address
	CallerAddress string `json:"callerAddress,omitempty"`

	// content length
	ContentLength int64 `json:"contentLength,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// content Uri
	// Format: uri
	ContentURI strfmt.URI `json:"contentUri,omitempty"`

	// created by
	CreatedBy *DomainEntityRef `json:"createdBy,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// download sharing Uri
	// Format: uri
	DownloadSharingURI strfmt.URI `json:"downloadSharingUri,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// page count
	PageCount int64 `json:"pageCount,omitempty"`

	// read
	Read bool `json:"read,omitempty"`

	// receiver address
	ReceiverAddress string `json:"receiverAddress,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// sharing Uri
	// Format: uri
	SharingURI strfmt.URI `json:"sharingUri,omitempty"`

	// thumbnails
	Thumbnails []*DocumentThumbnail `json:"thumbnails"`

	// workspace
	Workspace *DomainEntityRef `json:"workspace,omitempty"`
}

// Validate validates this fax document
func (m *FaxDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadSharingURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharingURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThumbnails(formats); err != nil {
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

func (m *FaxDocument) validateContentURI(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentURI) { // not required
		return nil
	}

	if err := validate.FormatOf("contentUri", "body", "uri", m.ContentURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *FaxDocument) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateDownloadSharingURI(formats strfmt.Registry) error {

	if swag.IsZero(m.DownloadSharingURI) { // not required
		return nil
	}

	if err := validate.FormatOf("downloadSharingUri", "body", "uri", m.DownloadSharingURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateSharingURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SharingURI) { // not required
		return nil
	}

	if err := validate.FormatOf("sharingUri", "body", "uri", m.SharingURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FaxDocument) validateThumbnails(formats strfmt.Registry) error {

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

func (m *FaxDocument) validateWorkspace(formats strfmt.Registry) error {

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
func (m *FaxDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FaxDocument) UnmarshalBinary(b []byte) error {
	var res FaxDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
