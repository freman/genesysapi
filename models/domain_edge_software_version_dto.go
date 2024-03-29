// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainEdgeSoftwareVersionDto domain edge software version dto
//
// swagger:model DomainEdgeSoftwareVersionDto
type DomainEdgeSoftwareVersionDto struct {

	// current
	Current bool `json:"current"`

	// edge Uri
	// Format: uri
	EdgeURI strfmt.URI `json:"edgeUri,omitempty"`

	// edge version
	EdgeVersion string `json:"edgeVersion,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// latest release
	LatestRelease bool `json:"latestRelease"`

	// name
	Name string `json:"name,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	PublishDate strfmt.DateTime `json:"publishDate,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this domain edge software version dto
func (m *DomainEdgeSoftwareVersionDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEdgeSoftwareVersionDto) validateEdgeURI(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeURI) { // not required
		return nil
	}

	if err := validate.FormatOf("edgeUri", "body", "uri", m.EdgeURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainEdgeSoftwareVersionDto) validatePublishDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishDate) { // not required
		return nil
	}

	if err := validate.FormatOf("publishDate", "body", "date-time", m.PublishDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainEdgeSoftwareVersionDto) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this domain edge software version dto based on the context it is used
func (m *DomainEdgeSoftwareVersionDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEdgeSoftwareVersionDto) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DomainEdgeSoftwareVersionDto) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainEdgeSoftwareVersionDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainEdgeSoftwareVersionDto) UnmarshalBinary(b []byte) error {
	var res DomainEdgeSoftwareVersionDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
