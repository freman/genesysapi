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

// ActionSurvey action survey
//
// swagger:model ActionSurvey
type ActionSurvey struct {

	// Questions shown to the user.
	// Required: true
	Questions []*JourneySurveyQuestion `json:"questions"`
}

// Validate validates this action survey
func (m *ActionSurvey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuestions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionSurvey) validateQuestions(formats strfmt.Registry) error {

	if err := validate.Required("questions", "body", m.Questions); err != nil {
		return err
	}

	for i := 0; i < len(m.Questions); i++ {
		if swag.IsZero(m.Questions[i]) { // not required
			continue
		}

		if m.Questions[i] != nil {
			if err := m.Questions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionSurvey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionSurvey) UnmarshalBinary(b []byte) error {
	var res ActionSurvey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}