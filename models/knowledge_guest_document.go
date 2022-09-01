// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KnowledgeGuestDocument knowledge guest document
//
// swagger:model KnowledgeGuestDocument
type KnowledgeGuestDocument struct {

	// List of alternate phrases related to the title which improves search results.
	Alternatives []*KnowledgeDocumentAlternative `json:"alternatives"`

	// The reference to category associated with the document.
	// Read Only: true
	Category *GuestCategoryReference `json:"category,omitempty"`

	// The user who created the document.
	// Read Only: true
	CreatedBy *UserReference `json:"createdBy,omitempty"`

	// Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The date on which the document was last published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DatePublished strfmt.DateTime `json:"datePublished,omitempty"`

	// The version of the document.
	DocumentVersion *AddressableEntityRef `json:"documentVersion,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The last published version number of the document.
	LastPublishedVersionNumber int32 `json:"lastPublishedVersionNumber,omitempty"`

	// The user who modified the document.
	// Read Only: true
	ModifiedBy *UserReference `json:"modifiedBy,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// ID of the guest session.
	// Read Only: true
	SessionID string `json:"sessionId,omitempty"`

	// State of the document.
	// Enum: [Draft Published Archived]
	State string `json:"state,omitempty"`

	// Document title.
	Title string `json:"title,omitempty"`

	// Variations of the document.
	Variations []*KnowledgeGuestDocumentVariation `json:"variations"`

	// Indicates if the knowledge document should be included in search results.
	Visible bool `json:"visible"`
}

// Validate validates this knowledge guest document
func (m *KnowledgeGuestDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternatives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
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

	if err := m.validateDatePublished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeGuestDocument) validateAlternatives(formats strfmt.Registry) error {

	if swag.IsZero(m.Alternatives) { // not required
		return nil
	}

	for i := 0; i < len(m.Alternatives); i++ {
		if swag.IsZero(m.Alternatives[i]) { // not required
			continue
		}

		if m.Alternatives[i] != nil {
			if err := m.Alternatives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alternatives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KnowledgeGuestDocument) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *KnowledgeGuestDocument) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateDatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.DatePublished) { // not required
		return nil
	}

	if err := validate.FormatOf("datePublished", "body", "date-time", m.DatePublished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateDocumentVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.DocumentVersion) { // not required
		return nil
	}

	if m.DocumentVersion != nil {
		if err := m.DocumentVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documentVersion")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var knowledgeGuestDocumentTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Draft","Published","Archived"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeGuestDocumentTypeStatePropEnum = append(knowledgeGuestDocumentTypeStatePropEnum, v)
	}
}

const (

	// KnowledgeGuestDocumentStateDraft captures enum value "Draft"
	KnowledgeGuestDocumentStateDraft string = "Draft"

	// KnowledgeGuestDocumentStatePublished captures enum value "Published"
	KnowledgeGuestDocumentStatePublished string = "Published"

	// KnowledgeGuestDocumentStateArchived captures enum value "Archived"
	KnowledgeGuestDocumentStateArchived string = "Archived"
)

// prop value enum
func (m *KnowledgeGuestDocument) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeGuestDocumentTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeGuestDocument) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeGuestDocument) validateVariations(formats strfmt.Registry) error {

	if swag.IsZero(m.Variations) { // not required
		return nil
	}

	for i := 0; i < len(m.Variations); i++ {
		if swag.IsZero(m.Variations[i]) { // not required
			continue
		}

		if m.Variations[i] != nil {
			if err := m.Variations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeGuestDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeGuestDocument) UnmarshalBinary(b []byte) error {
	var res KnowledgeGuestDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}