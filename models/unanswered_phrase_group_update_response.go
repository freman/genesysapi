// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UnansweredPhraseGroupUpdateResponse unanswered phrase group update response
//
// swagger:model UnansweredPhraseGroupUpdateResponse
type UnansweredPhraseGroupUpdateResponse struct {

	// Knowledge base unanswered group response
	Group *UnansweredGroup `json:"group,omitempty"`

	// List of phrases and documents linked in the patch request
	PhraseAssociations []*PhraseAssociations `json:"phraseAssociations"`
}

// Validate validates this unanswered phrase group update response
func (m *UnansweredPhraseGroupUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhraseAssociations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnansweredPhraseGroupUpdateResponse) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *UnansweredPhraseGroupUpdateResponse) validatePhraseAssociations(formats strfmt.Registry) error {
	if swag.IsZero(m.PhraseAssociations) { // not required
		return nil
	}

	for i := 0; i < len(m.PhraseAssociations); i++ {
		if swag.IsZero(m.PhraseAssociations[i]) { // not required
			continue
		}

		if m.PhraseAssociations[i] != nil {
			if err := m.PhraseAssociations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phraseAssociations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phraseAssociations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unanswered phrase group update response based on the context it is used
func (m *UnansweredPhraseGroupUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhraseAssociations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnansweredPhraseGroupUpdateResponse) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *UnansweredPhraseGroupUpdateResponse) contextValidatePhraseAssociations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhraseAssociations); i++ {

		if m.PhraseAssociations[i] != nil {
			if err := m.PhraseAssociations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phraseAssociations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phraseAssociations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnansweredPhraseGroupUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnansweredPhraseGroupUpdateResponse) UnmarshalBinary(b []byte) error {
	var res UnansweredPhraseGroupUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
