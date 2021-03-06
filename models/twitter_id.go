// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TwitterID User information for a twitter account
//
// swagger:model TwitterId
type TwitterID struct {

	// twitter user.id_str
	ID string `json:"id,omitempty"`

	// twitter user.name
	Name string `json:"name,omitempty"`

	// url of user's twitter profile
	// Read Only: true
	ProfileURL string `json:"profileUrl,omitempty"`

	// twitter user.screen_name
	ScreenName string `json:"screenName,omitempty"`

	// whether this data has been verified using the twitter API
	// Read Only: true
	Verified *bool `json:"verified"`
}

// Validate validates this twitter Id
func (m *TwitterID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TwitterID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TwitterID) UnmarshalBinary(b []byte) error {
	var res TwitterID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
