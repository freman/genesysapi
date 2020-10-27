// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DocumentationResult documentation result
//
// swagger:model DocumentationResult
type DocumentationResult struct {

	// The category of the documentation entity. Will be returned in responses for certain entities.
	Categories []int32 `json:"categories"`

	// The text or html content for the documentation entity. Will be returned in responses for certain entities.
	Content string `json:"content,omitempty"`

	// The description of the documentation entity. Will be returned in responses for certain entities.
	Description string `json:"description,omitempty"`

	// The excerpt of the documentation entity. Will be returned in responses for certain entities.
	Excerpt string `json:"excerpt,omitempty"`

	// The facet feature of the documentation entity. Will be returned in responses for certain entities.
	FacetFeature []int32 `json:"facet_feature"`

	// The facet role of the documentation entity. Will be returned in responses for certain entities.
	FacetRole []int32 `json:"facet_role"`

	// The facet service of the documentation entity. Will be returned in responses for certain entities.
	FacetService []int32 `json:"facet_service"`

	// The faq categories of the documentation entity. Will be returned in responses for certain entities.
	FaqCategories []int32 `json:"faq_categories"`

	// The search type. Will be returned in responses for certain entities.
	GetType string `json:"get_type,omitempty"`

	// The globally unique identifier for the object.
	// Required: true
	ID *int32 `json:"id"`

	// URL link for the documentation entity. Will be returned in responses for certain entities.
	Link string `json:"link,omitempty"`

	// The modified date for the documentation entity. Will be returned in responses for certain entities. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Modified strfmt.DateTime `json:"modified,omitempty"`

	// The name of the documentation entity. Will be returned in responses for certain entities.
	Name string `json:"name,omitempty"`

	// The releasenote category of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteCategory []int32 `json:"releasenote_category"`

	// The releasenote tag of the documentation entity. Will be returned in responses for certain entities.
	ReleasenoteTag []int32 `json:"releasenote_tag"`

	// The service of the documentation entity. Will be returned in responses for certain entities.
	Service []int32 `json:"service"`

	// The service area of the documentation entity. Will be returned in responses for certain entities.
	ServiceArea []int32 `json:"service-area"`

	// The slug of the documentation entity. Will be returned in responses for certain entities.
	Slug string `json:"slug,omitempty"`

	// The title of the documentation entity. Will be returned in responses for certain entities.
	Title string `json:"title,omitempty"`

	// The video categories of the documentation entity. Will be returned in responses for certain entities.
	VideoCategories []int32 `json:"video_categories"`
}

// Validate validates this documentation result
func (m *DocumentationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentationResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DocumentationResult) validateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.Modified) { // not required
		return nil
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocumentationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentationResult) UnmarshalBinary(b []byte) error {
	var res DocumentationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
