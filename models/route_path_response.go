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

// RoutePathResponse Route path configuration
//
// swagger:model RoutePathResponse
type RoutePathResponse struct {

	// The ID of the language associated with the route path
	Language *LanguageReference `json:"language,omitempty"`

	// The media type of the given queue associated with the route path
	// Enum: [Voice Chat Email Callback Message]
	MediaType string `json:"mediaType,omitempty"`

	// The ID of the queue associated with the route path
	Queue *QueueReference `json:"queue,omitempty"`

	// The set of skills associated with the route path
	// Unique: true
	Skills []*RoutingSkillReference `json:"skills"`
}

// Validate validates this route path response
func (m *RoutePathResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutePathResponse) validateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.Language) { // not required
		return nil
	}

	if m.Language != nil {
		if err := m.Language.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language")
			}
			return err
		}
	}

	return nil
}

var routePathResponseTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Voice","Chat","Email","Callback","Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routePathResponseTypeMediaTypePropEnum = append(routePathResponseTypeMediaTypePropEnum, v)
	}
}

const (

	// RoutePathResponseMediaTypeVoice captures enum value "Voice"
	RoutePathResponseMediaTypeVoice string = "Voice"

	// RoutePathResponseMediaTypeChat captures enum value "Chat"
	RoutePathResponseMediaTypeChat string = "Chat"

	// RoutePathResponseMediaTypeEmail captures enum value "Email"
	RoutePathResponseMediaTypeEmail string = "Email"

	// RoutePathResponseMediaTypeCallback captures enum value "Callback"
	RoutePathResponseMediaTypeCallback string = "Callback"

	// RoutePathResponseMediaTypeMessage captures enum value "Message"
	RoutePathResponseMediaTypeMessage string = "Message"
)

// prop value enum
func (m *RoutePathResponse) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routePathResponseTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RoutePathResponse) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

func (m *RoutePathResponse) validateQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *RoutePathResponse) validateSkills(formats strfmt.Registry) error {

	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	if err := validate.UniqueItems("skills", "body", m.Skills); err != nil {
		return err
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutePathResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutePathResponse) UnmarshalBinary(b []byte) error {
	var res RoutePathResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
